module github.com/ipfs-force-community/chain-co

go 1.16

require (
	github.com/dtynn/dix v0.1.0
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-bitfield v0.2.4
	github.com/filecoin-project/go-data-transfer v1.10.1
	github.com/filecoin-project/go-fil-markets v1.12.0
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.1-0.20210915140513-d354ccf10379
	github.com/filecoin-project/lotus v1.13.0-rc2
	github.com/filecoin-project/specs-actors v0.9.14
	github.com/filecoin-project/venus-auth v1.2.2-0.20210721103851-593a379c4916
	github.com/gbrlsnchs/jwt/v3 v3.0.0
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/ipfs-force-community/metrics v1.0.0
	github.com/ipfs-force-community/venus-common-utils v0.0.0-20210714031758-ea0e25ff0ec4
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-log/v2 v2.3.0
	github.com/ipfs/go-metrics-interface v0.0.1
	github.com/libp2p/go-libp2p-core v0.8.6
	github.com/urfave/cli/v2 v2.3.0
	github.com/whyrusleeping/pubsub v0.0.0-20190708150250-92bcb0691325
	go.opencensus.io v0.23.0
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.16.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/ipfs/go-ipfs-cmds => github.com/ipfs-force-community/go-ipfs-cmds v0.6.1-0.20210521090123-4587df7fa0ab

replace github.com/filecoin-project/go-jsonrpc => github.com/ipfs-force-community/go-jsonrpc v0.1.4-0.20211130034808-0af9ee3a3267

replace github.com/filecoin-project/lotus => github.com/ipfs-force-community/lotus v0.8.1-0.20211009103222-7455b2a31f37
