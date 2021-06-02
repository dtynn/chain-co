package proxy

import (
	"context"
	"github.com/dtynn/chain-co/api"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	api1 "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var _ ProxyAPI = (*Proxy)(nil)

type ProxyAPI interface {
	api.Proxy
}

type Proxy struct {
	Select func() (ProxyAPI, error)
}

// impl api.Proxy
func (p *Proxy) BeaconGetEntry(in0 context.Context, in1 abi.ChainEpoch) (out0 *types.BeaconEntry, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.BeaconGetEntry(in0, in1)
}

func (p *Proxy) ChainGetBlock(in0 context.Context, in1 cid.Cid) (out0 *types.BlockHeader, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetBlock(in0, in1)
}

func (p *Proxy) ChainGetBlockMessages(in0 context.Context, in1 cid.Cid) (out0 *api1.BlockMessages, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetBlockMessages(in0, in1)
}

func (p *Proxy) ChainGetGenesis(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetGenesis(in0)
}

func (p *Proxy) ChainGetMessage(in0 context.Context, in1 cid.Cid) (out0 *types.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetMessage(in0, in1)
}

func (p *Proxy) ChainGetParentMessages(in0 context.Context, in1 cid.Cid) (out0 []api1.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetParentMessages(in0, in1)
}

func (p *Proxy) ChainGetParentReceipts(in0 context.Context, in1 cid.Cid) (out0 []*types.MessageReceipt, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetParentReceipts(in0, in1)
}

func (p *Proxy) ChainGetRandomnessFromBeacon(in0 context.Context, in1 types.TipSetKey, in2 crypto.DomainSeparationTag, in3 abi.ChainEpoch, in4 []uint8) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetRandomnessFromBeacon(in0, in1, in2, in3, in4)
}

func (p *Proxy) ChainGetRandomnessFromTickets(in0 context.Context, in1 types.TipSetKey, in2 crypto.DomainSeparationTag, in3 abi.ChainEpoch, in4 []uint8) (out0 abi.Randomness, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetRandomnessFromTickets(in0, in1, in2, in3, in4)
}

func (p *Proxy) ChainGetTipSet(in0 context.Context, in1 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetTipSet(in0, in1)
}

func (p *Proxy) ChainGetTipSetByHeight(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetTipSetByHeight(in0, in1, in2)
}

func (p *Proxy) ChainHead(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainHead(in0)
}

func (p *Proxy) ChainTipSetWeight(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainTipSetWeight(in0, in1)
}
