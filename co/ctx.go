package co

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

// Ctx contains the shared components between different modules
type Ctx struct {
	lc      context.Context
	bcache  *blockHeaderCache
	headCh  chan *headCandidate
	nodeOpt NodeOption
}

type headCandidate struct {
	node   *Node
	ts     *types.TipSet
	weight types.BigInt
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
