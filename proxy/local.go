package proxy

import (
	"context"
	"fmt"

	api1 "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs-force-community/sophon-co/api"
)

var _ LocalAPI = (*Local)(nil)

type LocalAPI interface {
	api.Local
}

type Local struct {
	Select func(types.TipSetKey) (LocalAPI, error)
}

// impl api.Local
func (p *Local) ChainHead(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainHead %v", err)
		return
	}
	return cli.ChainHead(in0)
}

func (p *Local) ChainNotify(in0 context.Context) (out0 <-chan []*api1.HeadChange, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainNotify %v", err)
		return
	}
	return cli.ChainNotify(in0)
}
