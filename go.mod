module github.com/dtynn/chain-co

go 1.15

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-bitfield v0.2.4
	github.com/filecoin-project/go-data-transfer v1.2.7
	github.com/filecoin-project/go-fil-markets v1.1.9
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-multistore v0.0.3
	github.com/filecoin-project/go-state-types v0.1.0
	github.com/filecoin-project/lotus v1.8.0
	github.com/filecoin-project/specs-actors v0.9.13
	github.com/google/uuid v1.1.2
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-log/v2 v2.1.2-0.20200626104915-0016c0b4b3e4
	github.com/libp2p/go-libp2p-core v0.7.0
	github.com/whyrusleeping/pubsub v0.0.0-20190708150250-92bcb0691325
	go.uber.org/fx v1.9.0
	go.uber.org/zap v1.16.0
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
