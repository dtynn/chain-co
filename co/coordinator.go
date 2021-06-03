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

func NewCoordinator(ctx *Ctx, head *types.TipSet, weight types.BigInt, sel *Selector) (*Coordinator, error) {
	return &Coordinator{
		ctx:    ctx,
		head:   head,
		weight: weight,
		nodes:  make([]string, 0, 16),
		sel:    sel,
		tspub:  pubsub.New(256),
	}, nil
}

// Coordinator tries to setup the best nodes based on their incoming chain head
type Coordinator struct {
	ctx *Ctx

	headMu sync.RWMutex
	head   *types.TipSet
	weight types.BigInt
	nodes  []string

	sel *Selector

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

func (c *Coordinator) Stop() error {
	c.tspub.Shutdown()
	return nil
}

func (c *Coordinator) handleCandidate(hc *headCandidate) {
	clog := log.With("node", hc.node.info.Host, "h", hc.ts.Height(), "w", hc.weight, "drift", time.Now().Unix()-int64(hc.ts.MinTimestamp()))

	c.headMu.Lock()

	if c.head == nil || hc.weight.GreaterThan(c.weight) {
		clog.Debug("head replaced")

		prev := c.head
		next := hc.ts

		c.head = hc.ts
		c.weight = hc.weight
		c.nodes = append(c.nodes[:0], hc.node.info.Addr)
		c.sel.setPriors(hc.node.info.Addr)

		c.headMu.Unlock()

		if err := c.applyTipSetChange(prev, next, hc.node); err != nil {
			clog.Errorf("apply tipset change: %s", err)
		}

		return
	}

	if c.head.Equals(hc.ts) {
		contains := false
		for ni := range c.nodes {
			if c.nodes[ni] == hc.node.info.Addr {
				contains = true
				break
			}
		}

		if !contains {
			c.nodes = append(c.nodes, hc.node.info.Addr)
			c.sel.setPriors(c.nodes...)

			clog.Debug("another node caught up")
		}

		c.headMu.Unlock()
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
