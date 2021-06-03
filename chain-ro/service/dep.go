package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dtynn/dix"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"

	"github.com/dtynn/chain-co/co"
	"github.com/dtynn/chain-co/proxy"
)

const ExtractFullNodeAPIKey dix.Invoke = 1

func Build(ctx context.Context, overrides ...dix.Option) (dix.StopFunc, error) {
	opts := []dix.Option{
		dix.Override(new(co.NodeOption), co.DefaultNodeOption),
		dix.Override(new(*co.Ctx), co.NewCtx),
		dix.Override(new(*co.Connector), co.NewConnector),
		dix.Override(new(*co.Coordinator), buildCoordinator),
		dix.Override(new(*co.Selector), co.NewSelector),
		dix.Override(new(*proxy.Proxy), buildProxyAPI),
		dix.Override(new(*proxy.Local), buildLocalAPI),
		dix.Override(new(*proxy.UnSupport), buildUnSupportAPI),
	}
	opts = append(opts, overrides...)
	return dix.New(ctx, opts...)
}

func FullNode(full *api.FullNode) dix.Option {
	return dix.Override(ExtractFullNodeAPIKey, func(srv Service) error {
		*full = &srv
		return nil
	})
}

func ParseNodeInfoList(raws []string) dix.Option {
	return dix.Override(new(co.NodeInfoList), func() (co.NodeInfoList, error) {
		list := make(co.NodeInfoList, 0, len(raws))
		for _, str := range raws {
			info := co.ParseNodeInfo(str)
			if _, err := info.DialArgs(); err != nil {
				return nil, fmt.Errorf("invalid node info: %s", str)
			}

			list = append(list, info)
		}

		return list, nil
	})
}

func buildCoordinator(lc fx.Lifecycle, ctx *co.Ctx, connector *co.Connector, infos co.NodeInfoList, sel *co.Selector) (*co.Coordinator, error) {
	nodes := make([]*co.Node, 0, len(infos))
	allDone := false
	defer func() {
		if !allDone {
			for i := range nodes {
				nodes[i].Stop()
			}
		}
	}()

	var head *types.TipSet
	weight := types.NewInt(0)

	for i := range infos {
		info := infos[i]
		nlog := log.With("host", info.Host)

		node, err := connector.Connect(info)
		if err != nil {
			nlog.Errorf("connect failed: %s", err)
			continue
		}

		full := node.FullNode()
		h, w, err := getHeadCandidate(full)
		if err != nil {
			node.Stop()
			nlog.Errorf("failed to get head: %s", err)
			continue
		}

		if head == nil || w.GreaterThan(weight) {
			head = h
			weight = w
		}

		nodes = append(nodes, node)
	}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("no available node")
	}

	coordinator, err := co.NewCoordinator(ctx, head, weight, sel)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go coordinator.Start()
			sel.ReplaceNodes(nodes, nil, false)
			return nil
		},
		OnStop: func(context.Context) error {
			coordinator.Stop()
			return nil
		},
	})

	allDone = true
	return coordinator, nil
}

func getHeadCandidate(full api.FullNode) (*types.TipSet, types.BigInt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	head, err := full.ChainHead(ctx)
	if err != nil {
		return nil, types.BigInt{}, err
	}

	weight, err := full.ChainTipSetWeight(ctx, head.Key())
	if err != nil {
		return nil, types.BigInt{}, err
	}

	return head, weight, nil
}

func buildProxyAPI(sel *co.Selector) *proxy.Proxy {
	return &proxy.Proxy{
		Select: func() (proxy.ProxyAPI, error) {
			node, err := sel.Select()
			if err != nil {
				return nil, err
			}

			return node.FullNode(), nil
		},
	}
}

func buildLocalAPI(lsrv LocalChainService) *proxy.Local {
	return &proxy.Local{
		Select: func() (proxy.LocalAPI, error) {
			return &lsrv, nil
		},
	}
}

func buildUnSupportAPI() *proxy.UnSupport {
	return &proxy.UnSupport{
		Select: func() (proxy.UnSupportAPI, error) {
			return nil, fmt.Errorf("api not supported")
		},
	}
}
