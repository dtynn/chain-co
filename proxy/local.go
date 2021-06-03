package proxy

import (
	"context"
	"github.com/dtynn/chain-co/api"
	api1 "github.com/filecoin-project/lotus/api"
)

var _ LocalAPI = (*Local)(nil)

type LocalAPI interface {
	api.Local
}

type Local struct {
	Select func() (LocalAPI, error)
}

// impl api.Local
func (p *Local) ChainNotify(in0 context.Context) (out0 <-chan []*api1.HeadChange, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainNotify(in0)
}
