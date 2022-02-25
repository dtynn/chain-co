package service

import (
	logging "github.com/ipfs/go-log/v2"
	"go.uber.org/fx"

	"github.com/ipfs-force-community/chain-co/co"
	"github.com/ipfs-force-community/chain-co/proxy"
)

var log = logging.Logger("chain-ro-srv")

// LocalChainService impls proxy.Local
type LocalChainService struct {
	fx.In
	*co.Coordinator
}

// Service impls api.FullNode
type Service struct {
	fx.In

	*proxy.Proxy
	*proxy.Local
	*proxy.UnSupport
}
