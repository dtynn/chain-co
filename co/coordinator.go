package co

import (
	"fmt"
	"sync"
	"time"

	logging "github.com/ipfs/go-log/v2"
	"github.com/whyrusleeping/pubsub"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

// common errors
var (
	ErrNoNodeAvailable = fmt.Errorf("no node available")
)

var log = logging.Logger("chain-co")

const (
	tipsetChangeTopic = "tschange"
)

// func New(mctx helpers.MetricsCtx, lc fx.Lifecycle, nodes MasterNodeAddresses) (*Masters, error) {
//     if len(nodes) == 0 {
//         return nil, fmt.Errorf("no master node addrs provided")
//     }

//     allDone := false
//     clients := make(map[string]*Client)

//     lifeCtx := helpers.LifecycleCtx(mctx, lc)
//     runCtx, runCancel := context.WithCancel(lifeCtx)

//     defer func() {
//         if !allDone {
//             runCancel()
//             for _, cli := range clients {
//                 cli.Close(lifeCtx)
//             }
//         }
//     }()

//     var heaviest *types.TipSet
//     weight := types.NewInt(0)
//     prior := make([]string, 0, len(nodes))
//     all := make([]string, 0, len(nodes))

//     for ni := range nodes {
//         addr := nodes[ni]
//         cli, err := NewClient(runCtx, addr)
//         if err != nil {
//             return nil, fmt.Errorf("open master client for %s: %w", addr, err)
//         }

//         all = append(all, addr)
//         clients[addr] = cli

//         head, err := cli.full.ChainHead(runCtx)
//         if err != nil {
//             return nil, fmt.Errorf("get head from %s: %w", addr, err)
//         }

//         if head == nil {
//             cli.log.Warn("head not provided")
//             continue
//         }

//         if heaviest != nil && heaviest.Equals(head) {
//             prior = append(prior, addr)
//             continue
//         }

//         hw, err := cli.full.ChainTipSetWeight(runCtx, head.Key())
//         if err != nil {
//             return nil, fmt.Errorf("get weight from %s: %w", addr, err)
//         }

//         if hw.GreaterThan(weight) {
//             heaviest = head
//             weight = hw
//             prior = append(prior[:0], addr)
//         }
//     }

//     if heaviest == nil {
//         return nil, fmt.Errorf("unable to get heaviest chain head from masters")
//     }

//     allDone = true
//     m := &Masters{
//         ctx:    runCtx,
//         cancel: runCancel,
//     }

//     m.prior.addrs = prior

//     m.all.addrs = all
//     m.all.clients = clients

//     m.head.ch = make(chan HeadChange, 16)
//     m.head.ts = heaviest
//     m.head.weight = weight
//     m.head.nodes = make([]string, len(prior))
//     copy(m.head.nodes, prior)

//     return m, nil
// }

// Coordinator tries to setup the best nodes based on their incoming chain head
type Coordinator struct {
	ctx *Ctx

	mu     sync.Mutex
	head   *types.TipSet
	weight types.BigInt
	nodes  []string

	sel   *Selector
	tspub *pubsub.PubSub
}

// Start starts the coordinate loop
func (c *Coordinator) Start() {
	log.Info("start head coordinator loop")
	defer log.Info("stop head coordinator loop")

	for {
		select {
		case <-c.ctx.lc.Done():
			return

		case hc := <-c.ctx.headCh:
			c.handleCandidate(hc)
		}
	}
}

func (c *Coordinator) handleCandidate(hc *headCandidate) {
	clog := log.With("node", hc.node.info.Host, "h", hc.ts.Height(), "w", hc.weight, "drift", time.Now().Unix()-int64(hc.ts.MinTimestamp()))
	if c.head == nil || hc.weight.GreaterThan(c.weight) {
		clog.Debug("head replaced")

		c.mu.Lock()
		prev := c.head
		next := hc.ts

		c.head = hc.ts
		c.weight = hc.weight
		c.nodes = append(c.nodes[:0], hc.node.info.Host)
		c.mu.Unlock()

		c.sel.setPriors(hc.node.info.Host)
		if err := c.applyTipSetChange(prev, next, hc.node); err != nil {
			clog.Errorf("apply tipset change: %s", err)
		}

		return
	}

	if c.head.Equals(hc.ts) {
		contains := false
		for ni := range c.nodes {
			if c.nodes[ni] == hc.node.info.Host {
				contains = true
				break
			}
		}

		if !contains {
			c.mu.Lock()
			c.nodes = append(c.nodes, hc.node.info.Host)
			c.mu.Unlock()

			c.sel.setPriors(c.nodes...)

			clog.Debug("another node caught up")
		}

		return
	}

	clog.Debug("ignored a lighter head")
	return
}

func (c *Coordinator) applyTipSetChange(prev, next *types.TipSet, node *Node) error {
	revert, apply, err := store.ReorgOps(node.loadTipSet, prev, next)
	if err != nil {
		return err
	}

	hc := make([]*api.HeadChange, 0, len(revert)+len(apply))
	for i := range revert {
		hc = append(hc, &api.HeadChange{
			Type: store.HCRevert,
			Val:  revert[i],
		})
	}

	for i := range apply {
		hc = append(hc, &api.HeadChange{
			Type: store.HCApply,
			Val:  apply[i],
		})
	}

	if len(hc) == 0 {
		return nil
	}

	c.tspub.Pub(hc, tipsetChangeTopic)
	return nil
}
