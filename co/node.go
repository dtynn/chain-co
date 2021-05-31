package co

import (
	"context"
	"net/http"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/hashicorp/go-multierror"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

const (
	rpcV0Endpoint = "/rpc/v0"
)

// NodeOption is for node configuration
type NodeOption struct {
	ReListenMinInterval time.Duration
	ReListenMaxInterval time.Duration

	APITimeout time.Duration
}

// NodeInfo for connection
type NodeInfo struct {
	Host   string
	Header http.Header
	UseTLS bool
}

func connect(sctx *Ctx, info NodeInfo) (*Node, error) {
	scheme := "ws://"
	if info.UseTLS {
		scheme = "wss://"
	}

	addr := scheme + info.Host + rpcV0Endpoint

	full, closer, err := client.NewFullNodeRPC(sctx.lc, addr, info.Header)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(sctx.lc)
	node := &Node{
		opt:    sctx.nodeOpt,
		info:   info,
		ctx:    ctx,
		cancel: cancel,
		sctx:   sctx,
		log:    log.With("remote", addr),
	}

	node.upstream.full = full
	node.upstream.closer = closer

	return node, nil
}

// Node is a FullNode client
type Node struct {
	opt  NodeOption
	info NodeInfo

	reListenInterval time.Duration

	ctx    context.Context
	cancel context.CancelFunc

	sctx *Ctx

	upstream struct {
		full   api.FullNode
		closer jsonrpc.ClientCloser
	}

	log *zap.SugaredLogger
}

// Start starts a head change loop
func (n *Node) Start() {
	n.log.Info("start head change loop")
	defer n.log.Info("stop head change loop")

	for {
		ch, err := n.reListen()
		if err != nil {
			n.log.Errorf("failed to listen head change: %s", err)
			return
		}

		chLifeCtx, chLifeCancel := context.WithCancel(n.ctx)

	CHANGES_LOOP:
		for {
			select {
			case <-n.ctx.Done():
				chLifeCancel()
				return

			case changes, ok := <-ch:
				if !ok {
					break CHANGES_LOOP
				}

				go n.applyChanges(chLifeCtx, changes)
			}
		}

		chLifeCancel()
	}
}

// Stop closes current node
func (n *Node) Stop() error {
	n.cancel()
	n.upstream.closer()
	return nil
}

func (n *Node) reListen() (<-chan []*api.HeadChange, error) {
	for {
		ch, err := n.upstream.full.ChainNotify(n.ctx)
		if err != nil {
			n.log.Errorf("call CahinNotify: %s, will re-call in %s", err, n.reListenInterval)

			select {
			case <-n.ctx.Done():
				return nil, n.ctx.Err()

			case <-time.After(n.reListenInterval):

				n.reListenInterval *= 2
				if n.reListenInterval > n.opt.ReListenMaxInterval {
					n.reListenInterval = n.opt.ReListenMaxInterval
				}

			}

			continue
		}

		n.reListenInterval = n.opt.ReListenMinInterval
		return ch, nil
	}
}

func (n *Node) applyChanges(lifeCtx context.Context, changes []*api.HeadChange) {
	n.sctx.bcache.add(changes)

	idx := -1
	for i := range changes {
		switch changes[i].Type {
		case store.HCCurrent, store.HCApply:
			idx = i
		}
	}

	if idx == -1 {
		return
	}

	ts := changes[idx].Val

	callCtx, callCancel := context.WithTimeout(lifeCtx, n.opt.APITimeout)
	weight, err := n.upstream.full.ChainTipSetWeight(callCtx, ts.Key())
	callCancel()

	if err != nil {
		n.log.Errorf("call ChainTipSetWeight: %s", err)
		return
	}

	hc := &headCandidate{
		node:   n,
		ts:     ts,
		weight: weight,
	}

	slow := time.NewTicker(5 * time.Second)
	defer slow.Stop()

	t := time.Now()

	for {
		select {
		case <-lifeCtx.Done():
			return

		case n.sctx.headCh <- hc:
			return

		case tick := <-slow.C:
			n.log.Warnf("it took too long before we can send the new head change, ts=%s, height=%d, weight=%s, delay=%s", ts.Key(), ts.Height(), weight, tick.Sub(t))
		}
	}
}

func (n *Node) loadTipSet(tsk types.TipSetKey) (*types.TipSet, error) {
	reqCtx, reqCancel := context.WithTimeout(n.ctx, n.opt.APITimeout)
	defer reqCancel()

	var wg multierror.Group
	cids := tsk.Cids()
	blks := make([]*types.BlockHeader, len(cids))
	for ci := range cids {
		i := ci
		wg.Go(func() error {
			blk, err := n.loadBlockHeader(reqCtx, cids[i])
			if err != nil {
				return err
			}

			blks[i] = blk
			return nil
		})
	}

	if err := wg.Wait(); err != nil {
		return nil, err
	}

	return types.NewTipSet(blks)
}

func (n *Node) loadBlockHeader(ctx context.Context, c cid.Cid) (*types.BlockHeader, error) {
	if blk, ok := n.sctx.bcache.load(c); ok {
		return blk, nil
	}

	blk, err := n.upstream.full.ChainGetBlock(ctx, c)
	return blk, err
}
