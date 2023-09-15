package co

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ipfs/go-cid"
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

// NewCtx constructs a Ctx instance
func NewCtx(mctx helpers.MetricsCtx, lc fx.Lifecycle, nodeOpt NodeOption) (*Ctx, error) {
	return &Ctx{
		lc:        helpers.LifecycleCtx(mctx, lc),
		headCh:    make(chan *headCandidate, 256),
		errNodeCh: make(chan string, 256),
		nodeOpt:   nodeOpt,
	}, nil
}

// Ctx contains the shared components between different modules
type Ctx struct {
	lc        context.Context
	headCh    chan *headCandidate
	errNodeCh chan string

	nodeOpt NodeOption
}

type headCandidate struct {
	node   *Node
	ts     *types.TipSet
	weight types.BigInt
}

func newBlockHeaderCache(size int) (*blockHeaderCache, error) {
	cache, err := lru.New2Q(size)
	if err != nil {
		return nil, err
	}

	return &blockHeaderCache{
		cache: cache,
	}, nil
}

type blockHeaderCache struct {
	cache *lru.TwoQueueCache
}

func (bc *blockHeaderCache) add(changes []*api.HeadChange) {
	for _, hc := range changes {
		blks := hc.Val.Blocks()
		for bi := range blks {
			bc.cache.Add(blks[bi].Cid(), blks[bi])
		}
	}
}

func (bc *blockHeaderCache) load(c cid.Cid) (*types.BlockHeader, bool) {
	val, ok := bc.cache.Get(c)
	if !ok {
		return nil, false
	}

	blk, ok := val.(*types.BlockHeader)
	return blk, ok
}

func (bc *blockHeaderCache) has(c cid.Cid) bool {
	_, ok := bc.cache.Peek(c)
	return ok
}

func (bc *blockHeaderCache) hasKey(key types.TipSetKey) bool {
	for _, blkCid := range key.Cids() {
		has := bc.has(blkCid)
		if !has {
			return false
		}
	}
	return true
}
