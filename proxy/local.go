package proxy

import (
	"context"
	api1 "github.com/filecoin-project/lotus/api"
	"github.com/ipfs-force-community/chain-co/api"
	"golang.org/x/xerrors"
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
		err = xerrors.Errorf("api ChainNotify %v", err)
		return
	}
	return cli.ChainNotify(in0)
}
