package proxy

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/go-state-types/network"
	api1 "github.com/filecoin-project/lotus/api"
	miner1 "github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs-force-community/chain-co/api"
	"github.com/ipfs/go-cid"
)

var _ ProxyAPI = (*Proxy)(nil)

type ProxyAPI interface {
	api.Proxy
}

type Proxy struct {
	Select func(types.TipSetKey) (ProxyAPI, error)
}

// impl api.Proxy
func (p *Proxy) ChainGetBlock(in0 context.Context, in1 cid.Cid) (out0 *types.BlockHeader, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetBlock %v", err)
		return
	}
	return cli.ChainGetBlock(in0, in1)
}

func (p *Proxy) ChainGetBlockMessages(in0 context.Context, in1 cid.Cid) (out0 *api1.BlockMessages, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetBlockMessages %v", err)
		return
	}
	return cli.ChainGetBlockMessages(in0, in1)
}

func (p *Proxy) ChainGetEvents(in0 context.Context, in1 cid.Cid) (out0 []types.Event, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetEvents %v", err)
		return
	}
	return cli.ChainGetEvents(in0, in1)
}

func (p *Proxy) ChainGetGenesis(in0 context.Context) (out0 *types.TipSet, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetGenesis %v", err)
		return
	}
	return cli.ChainGetGenesis(in0)
}

func (p *Proxy) ChainGetMessage(in0 context.Context, in1 cid.Cid) (out0 *types.Message, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetMessage %v", err)
		return
	}
	return cli.ChainGetMessage(in0, in1)
}

func (p *Proxy) ChainGetMessagesInTipset(in0 context.Context, in1 types.TipSetKey) (out0 []api1.Message, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api ChainGetMessagesInTipset %v", err)
		return
	}
	return cli.ChainGetMessagesInTipset(in0, in1)
}

func (p *Proxy) ChainGetParentMessages(in0 context.Context, in1 cid.Cid) (out0 []api1.Message, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetParentMessages %v", err)
		return
	}
	return cli.ChainGetParentMessages(in0, in1)
}

func (p *Proxy) ChainGetParentReceipts(in0 context.Context, in1 cid.Cid) (out0 []*types.MessageReceipt, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainGetParentReceipts %v", err)
		return
	}
	return cli.ChainGetParentReceipts(in0, in1)
}

func (p *Proxy) ChainGetPath(in0 context.Context, in1 types.TipSetKey, in2 types.TipSetKey) (out0 []*api1.HeadChange, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api ChainGetPath %v", err)
		return
	}
	return cli.ChainGetPath(in0, in1, in2)
}

func (p *Proxy) ChainGetTipSet(in0 context.Context, in1 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api ChainGetTipSet %v", err)
		return
	}
	return cli.ChainGetTipSet(in0, in1)
}

func (p *Proxy) ChainGetTipSetAfterHeight(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api ChainGetTipSetAfterHeight %v", err)
		return
	}
	return cli.ChainGetTipSetAfterHeight(in0, in1, in2)
}

func (p *Proxy) ChainGetTipSetByHeight(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 *types.TipSet, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api ChainGetTipSetByHeight %v", err)
		return
	}
	return cli.ChainGetTipSetByHeight(in0, in1, in2)
}

func (p *Proxy) ChainHasObj(in0 context.Context, in1 cid.Cid) (out0 bool, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainHasObj %v", err)
		return
	}
	return cli.ChainHasObj(in0, in1)
}

func (p *Proxy) ChainReadObj(in0 context.Context, in1 cid.Cid) (out0 []uint8, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainReadObj %v", err)
		return
	}
	return cli.ChainReadObj(in0, in1)
}

func (p *Proxy) ChainStatObj(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 api1.ObjStat, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api ChainStatObj %v", err)
		return
	}
	return cli.ChainStatObj(in0, in1, in2)
}

func (p *Proxy) ChainTipSetWeight(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api ChainTipSetWeight %v", err)
		return
	}
	return cli.ChainTipSetWeight(in0, in1)
}

func (p *Proxy) EthAccounts(in0 context.Context) (out0 []ethtypes.EthAddress, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthAccounts %v", err)
		return
	}
	return cli.EthAccounts(in0)
}

func (p *Proxy) EthAddressToFilecoinAddress(in0 context.Context, in1 ethtypes.EthAddress) (out0 address.Address, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthAddressToFilecoinAddress %v", err)
		return
	}
	return cli.EthAddressToFilecoinAddress(in0, in1)
}

func (p *Proxy) EthBlockNumber(in0 context.Context) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthBlockNumber %v", err)
		return
	}
	return cli.EthBlockNumber(in0)
}

func (p *Proxy) EthCall(in0 context.Context, in1 ethtypes.EthCall, in2 string) (out0 ethtypes.EthBytes, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthCall %v", err)
		return
	}
	return cli.EthCall(in0, in1, in2)
}

func (p *Proxy) EthChainId(in0 context.Context) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthChainId %v", err)
		return
	}
	return cli.EthChainId(in0)
}

func (p *Proxy) EthEstimateGas(in0 context.Context, in1 ethtypes.EthCall) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthEstimateGas %v", err)
		return
	}
	return cli.EthEstimateGas(in0, in1)
}

func (p *Proxy) EthFeeHistory(in0 context.Context, in1 jsonrpc.RawParams) (out0 ethtypes.EthFeeHistory, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthFeeHistory %v", err)
		return
	}
	return cli.EthFeeHistory(in0, in1)
}

func (p *Proxy) EthGasPrice(in0 context.Context) (out0 ethtypes.EthBigInt, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGasPrice %v", err)
		return
	}
	return cli.EthGasPrice(in0)
}

func (p *Proxy) EthGetBalance(in0 context.Context, in1 ethtypes.EthAddress, in2 string) (out0 ethtypes.EthBigInt, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetBalance %v", err)
		return
	}
	return cli.EthGetBalance(in0, in1, in2)
}

func (p *Proxy) EthGetBlockByHash(in0 context.Context, in1 ethtypes.EthHash, in2 bool) (out0 ethtypes.EthBlock, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetBlockByHash %v", err)
		return
	}
	return cli.EthGetBlockByHash(in0, in1, in2)
}

func (p *Proxy) EthGetBlockByNumber(in0 context.Context, in1 string, in2 bool) (out0 ethtypes.EthBlock, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetBlockByNumber %v", err)
		return
	}
	return cli.EthGetBlockByNumber(in0, in1, in2)
}

func (p *Proxy) EthGetBlockTransactionCountByHash(in0 context.Context, in1 ethtypes.EthHash) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetBlockTransactionCountByHash %v", err)
		return
	}
	return cli.EthGetBlockTransactionCountByHash(in0, in1)
}

func (p *Proxy) EthGetBlockTransactionCountByNumber(in0 context.Context, in1 ethtypes.EthUint64) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetBlockTransactionCountByNumber %v", err)
		return
	}
	return cli.EthGetBlockTransactionCountByNumber(in0, in1)
}

func (p *Proxy) EthGetCode(in0 context.Context, in1 ethtypes.EthAddress, in2 string) (out0 ethtypes.EthBytes, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetCode %v", err)
		return
	}
	return cli.EthGetCode(in0, in1, in2)
}

func (p *Proxy) EthGetFilterChanges(in0 context.Context, in1 ethtypes.EthFilterID) (out0 *ethtypes.EthFilterResult, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetFilterChanges %v", err)
		return
	}
	return cli.EthGetFilterChanges(in0, in1)
}

func (p *Proxy) EthGetFilterLogs(in0 context.Context, in1 ethtypes.EthFilterID) (out0 *ethtypes.EthFilterResult, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetFilterLogs %v", err)
		return
	}
	return cli.EthGetFilterLogs(in0, in1)
}

func (p *Proxy) EthGetLogs(in0 context.Context, in1 *ethtypes.EthFilterSpec) (out0 *ethtypes.EthFilterResult, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetLogs %v", err)
		return
	}
	return cli.EthGetLogs(in0, in1)
}

func (p *Proxy) EthGetMessageCidByTransactionHash(in0 context.Context, in1 *ethtypes.EthHash) (out0 *cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetMessageCidByTransactionHash %v", err)
		return
	}
	return cli.EthGetMessageCidByTransactionHash(in0, in1)
}

func (p *Proxy) EthGetStorageAt(in0 context.Context, in1 ethtypes.EthAddress, in2 ethtypes.EthBytes, in3 string) (out0 ethtypes.EthBytes, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetStorageAt %v", err)
		return
	}
	return cli.EthGetStorageAt(in0, in1, in2, in3)
}

func (p *Proxy) EthGetTransactionByBlockHashAndIndex(in0 context.Context, in1 ethtypes.EthHash, in2 ethtypes.EthUint64) (out0 ethtypes.EthTx, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionByBlockHashAndIndex %v", err)
		return
	}
	return cli.EthGetTransactionByBlockHashAndIndex(in0, in1, in2)
}

func (p *Proxy) EthGetTransactionByBlockNumberAndIndex(in0 context.Context, in1 ethtypes.EthUint64, in2 ethtypes.EthUint64) (out0 ethtypes.EthTx, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionByBlockNumberAndIndex %v", err)
		return
	}
	return cli.EthGetTransactionByBlockNumberAndIndex(in0, in1, in2)
}

func (p *Proxy) EthGetTransactionByHash(in0 context.Context, in1 *ethtypes.EthHash) (out0 *ethtypes.EthTx, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionByHash %v", err)
		return
	}
	return cli.EthGetTransactionByHash(in0, in1)
}

func (p *Proxy) EthGetTransactionCount(in0 context.Context, in1 ethtypes.EthAddress, in2 string) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionCount %v", err)
		return
	}
	return cli.EthGetTransactionCount(in0, in1, in2)
}

func (p *Proxy) EthGetTransactionHashByCid(in0 context.Context, in1 cid.Cid) (out0 *ethtypes.EthHash, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionHashByCid %v", err)
		return
	}
	return cli.EthGetTransactionHashByCid(in0, in1)
}

func (p *Proxy) EthGetTransactionReceipt(in0 context.Context, in1 ethtypes.EthHash) (out0 *api1.EthTxReceipt, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthGetTransactionReceipt %v", err)
		return
	}
	return cli.EthGetTransactionReceipt(in0, in1)
}

func (p *Proxy) EthMaxPriorityFeePerGas(in0 context.Context) (out0 ethtypes.EthBigInt, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthMaxPriorityFeePerGas %v", err)
		return
	}
	return cli.EthMaxPriorityFeePerGas(in0)
}

func (p *Proxy) EthNewBlockFilter(in0 context.Context) (out0 ethtypes.EthFilterID, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthNewBlockFilter %v", err)
		return
	}
	return cli.EthNewBlockFilter(in0)
}

func (p *Proxy) EthNewFilter(in0 context.Context, in1 *ethtypes.EthFilterSpec) (out0 ethtypes.EthFilterID, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthNewFilter %v", err)
		return
	}
	return cli.EthNewFilter(in0, in1)
}

func (p *Proxy) EthNewPendingTransactionFilter(in0 context.Context) (out0 ethtypes.EthFilterID, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthNewPendingTransactionFilter %v", err)
		return
	}
	return cli.EthNewPendingTransactionFilter(in0)
}

func (p *Proxy) EthProtocolVersion(in0 context.Context) (out0 ethtypes.EthUint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthProtocolVersion %v", err)
		return
	}
	return cli.EthProtocolVersion(in0)
}

func (p *Proxy) EthSendRawTransaction(in0 context.Context, in1 ethtypes.EthBytes) (out0 ethtypes.EthHash, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthSendRawTransaction %v", err)
		return
	}
	return cli.EthSendRawTransaction(in0, in1)
}

func (p *Proxy) EthSubscribe(in0 context.Context, in1 jsonrpc.RawParams) (out0 ethtypes.EthSubscriptionID, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthSubscribe %v", err)
		return
	}
	return cli.EthSubscribe(in0, in1)
}

func (p *Proxy) EthUninstallFilter(in0 context.Context, in1 ethtypes.EthFilterID) (out0 bool, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthUninstallFilter %v", err)
		return
	}
	return cli.EthUninstallFilter(in0, in1)
}

func (p *Proxy) EthUnsubscribe(in0 context.Context, in1 ethtypes.EthSubscriptionID) (out0 bool, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api EthUnsubscribe %v", err)
		return
	}
	return cli.EthUnsubscribe(in0, in1)
}

func (p *Proxy) FilecoinAddressToEthAddress(in0 context.Context, in1 address.Address) (out0 ethtypes.EthAddress, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api FilecoinAddressToEthAddress %v", err)
		return
	}
	return cli.FilecoinAddressToEthAddress(in0, in1)
}

func (p *Proxy) GasBatchEstimateMessageGas(in0 context.Context, in1 []*api1.EstimateMessage, in2 uint64, in3 types.TipSetKey) (out0 []*api1.EstimateResult, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api GasBatchEstimateMessageGas %v", err)
		return
	}
	return cli.GasBatchEstimateMessageGas(in0, in1, in2, in3)
}

func (p *Proxy) GasEstimateFeeCap(in0 context.Context, in1 *types.Message, in2 int64, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api GasEstimateFeeCap %v", err)
		return
	}
	return cli.GasEstimateFeeCap(in0, in1, in2, in3)
}

func (p *Proxy) GasEstimateGasLimit(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 int64, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api GasEstimateGasLimit %v", err)
		return
	}
	return cli.GasEstimateGasLimit(in0, in1, in2)
}

func (p *Proxy) GasEstimateGasPremium(in0 context.Context, in1 uint64, in2 address.Address, in3 int64, in4 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in4)
	if err != nil {
		err = fmt.Errorf("api GasEstimateGasPremium %v", err)
		return
	}
	return cli.GasEstimateGasPremium(in0, in1, in2, in3, in4)
}

func (p *Proxy) GasEstimateMessageGas(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec, in3 types.TipSetKey) (out0 *types.Message, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api GasEstimateMessageGas %v", err)
		return
	}
	return cli.GasEstimateMessageGas(in0, in1, in2, in3)
}

func (p *Proxy) MinerCreateBlock(in0 context.Context, in1 *api1.BlockTemplate) (out0 *types.BlockMsg, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MinerCreateBlock %v", err)
		return
	}
	return cli.MinerCreateBlock(in0, in1)
}

func (p *Proxy) MinerGetBaseInfo(in0 context.Context, in1 address.Address, in2 abi.ChainEpoch, in3 types.TipSetKey) (out0 *api1.MiningBaseInfo, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api MinerGetBaseInfo %v", err)
		return
	}
	return cli.MinerGetBaseInfo(in0, in1, in2, in3)
}

func (p *Proxy) MpoolBatchPush(in0 context.Context, in1 []*types.SignedMessage) (out0 []cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolBatchPush %v", err)
		return
	}
	return cli.MpoolBatchPush(in0, in1)
}

func (p *Proxy) MpoolBatchPushUntrusted(in0 context.Context, in1 []*types.SignedMessage) (out0 []cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolBatchPushUntrusted %v", err)
		return
	}
	return cli.MpoolBatchPushUntrusted(in0, in1)
}

func (p *Proxy) MpoolGetConfig(in0 context.Context) (out0 *types.MpoolConfig, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolGetConfig %v", err)
		return
	}
	return cli.MpoolGetConfig(in0)
}

func (p *Proxy) MpoolGetNonce(in0 context.Context, in1 address.Address) (out0 uint64, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolGetNonce %v", err)
		return
	}
	return cli.MpoolGetNonce(in0, in1)
}

func (p *Proxy) MpoolPending(in0 context.Context, in1 types.TipSetKey) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api MpoolPending %v", err)
		return
	}
	return cli.MpoolPending(in0, in1)
}

func (p *Proxy) MpoolPublishByAddr(in0 context.Context, in1 address.Address) (err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolPublishByAddr %v", err)
		return
	}
	return cli.MpoolPublishByAddr(in0, in1)
}

func (p *Proxy) MpoolPublishMessage(in0 context.Context, in1 *types.SignedMessage) (err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolPublishMessage %v", err)
		return
	}
	return cli.MpoolPublishMessage(in0, in1)
}

func (p *Proxy) MpoolPush(in0 context.Context, in1 *types.SignedMessage) (out0 cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolPush %v", err)
		return
	}
	return cli.MpoolPush(in0, in1)
}

func (p *Proxy) MpoolPushMessage(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec) (out0 *types.SignedMessage, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolPushMessage %v", err)
		return
	}
	return cli.MpoolPushMessage(in0, in1, in2)
}

func (p *Proxy) MpoolSelect(in0 context.Context, in1 types.TipSetKey, in2 float64) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolSelect %v", err)
		return
	}
	return cli.MpoolSelect(in0, in1, in2)
}

func (p *Proxy) MpoolSelects(in0 context.Context, in1 types.TipSetKey, in2 []float64) (out0 [][]*types.SignedMessage, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api MpoolSelects %v", err)
		return
	}
	return cli.MpoolSelects(in0, in1, in2)
}

func (p *Proxy) NetListening(in0 context.Context) (out0 bool, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api NetListening %v", err)
		return
	}
	return cli.NetListening(in0)
}

func (p *Proxy) NetVersion(in0 context.Context) (out0 string, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api NetVersion %v", err)
		return
	}
	return cli.NetVersion(in0)
}

func (p *Proxy) StateAccountKey(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateAccountKey %v", err)
		return
	}
	return cli.StateAccountKey(in0, in1, in2)
}

func (p *Proxy) StateActorCodeCIDs(in0 context.Context, in1 network.Version) (out0 map[string]cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateActorCodeCIDs %v", err)
		return
	}
	return cli.StateActorCodeCIDs(in0, in1)
}

func (p *Proxy) StateActorManifestCID(in0 context.Context, in1 network.Version) (out0 cid.Cid, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateActorManifestCID %v", err)
		return
	}
	return cli.StateActorManifestCID(in0, in1)
}

func (p *Proxy) StateAllMinerFaults(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 []*api1.Fault, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateAllMinerFaults %v", err)
		return
	}
	return cli.StateAllMinerFaults(in0, in1, in2)
}

func (p *Proxy) StateCall(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 *api1.InvocResult, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateCall %v", err)
		return
	}
	return cli.StateCall(in0, in1, in2)
}

func (p *Proxy) StateChangedActors(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 map[string]types.ActorV5, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateChangedActors %v", err)
		return
	}
	return cli.StateChangedActors(in0, in1, in2)
}

func (p *Proxy) StateCirculatingSupply(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateCirculatingSupply %v", err)
		return
	}
	return cli.StateCirculatingSupply(in0, in1)
}

func (p *Proxy) StateComputeDataCID(in0 context.Context, in1 address.Address, in2 abi.RegisteredSealProof, in3 []abi.DealID, in4 types.TipSetKey) (out0 cid.Cid, err error) {
	cli, err := p.Select(in4)
	if err != nil {
		err = fmt.Errorf("api StateComputeDataCID %v", err)
		return
	}
	return cli.StateComputeDataCID(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateDealProviderCollateralBounds(in0 context.Context, in1 abi.PaddedPieceSize, in2 bool, in3 types.TipSetKey) (out0 api1.DealCollateralBounds, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateDealProviderCollateralBounds %v", err)
		return
	}
	return cli.StateDealProviderCollateralBounds(in0, in1, in2, in3)
}

func (p *Proxy) StateGetActor(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *types.ActorV5, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateGetActor %v", err)
		return
	}
	return cli.StateGetActor(in0, in1, in2)
}

func (p *Proxy) StateGetAllocation(in0 context.Context, in1 address.Address, in2 verifreg.AllocationId, in3 types.TipSetKey) (out0 *verifreg.Allocation, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateGetAllocation %v", err)
		return
	}
	return cli.StateGetAllocation(in0, in1, in2, in3)
}

func (p *Proxy) StateGetAllocationForPendingDeal(in0 context.Context, in1 abi.DealID, in2 types.TipSetKey) (out0 *verifreg.Allocation, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateGetAllocationForPendingDeal %v", err)
		return
	}
	return cli.StateGetAllocationForPendingDeal(in0, in1, in2)
}

func (p *Proxy) StateGetAllocations(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 map[verifreg.AllocationId]verifreg.Allocation, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateGetAllocations %v", err)
		return
	}
	return cli.StateGetAllocations(in0, in1, in2)
}

func (p *Proxy) StateGetBeaconEntry(in0 context.Context, in1 abi.ChainEpoch) (out0 *types.BeaconEntry, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateGetBeaconEntry %v", err)
		return
	}
	return cli.StateGetBeaconEntry(in0, in1)
}

func (p *Proxy) StateGetClaim(in0 context.Context, in1 address.Address, in2 verifreg.ClaimId, in3 types.TipSetKey) (out0 *verifreg.Claim, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateGetClaim %v", err)
		return
	}
	return cli.StateGetClaim(in0, in1, in2, in3)
}

func (p *Proxy) StateGetClaims(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 map[verifreg.ClaimId]verifreg.Claim, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateGetClaims %v", err)
		return
	}
	return cli.StateGetClaims(in0, in1, in2)
}

func (p *Proxy) StateGetNetworkParams(in0 context.Context) (out0 *api1.NetworkParams, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateGetNetworkParams %v", err)
		return
	}
	return cli.StateGetNetworkParams(in0)
}

func (p *Proxy) StateGetRandomnessFromBeacon(in0 context.Context, in1 crypto.DomainSeparationTag, in2 abi.ChainEpoch, in3 []uint8, in4 types.TipSetKey) (out0 abi.Randomness, err error) {
	cli, err := p.Select(in4)
	if err != nil {
		err = fmt.Errorf("api StateGetRandomnessFromBeacon %v", err)
		return
	}
	return cli.StateGetRandomnessFromBeacon(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateGetRandomnessFromTickets(in0 context.Context, in1 crypto.DomainSeparationTag, in2 abi.ChainEpoch, in3 []uint8, in4 types.TipSetKey) (out0 abi.Randomness, err error) {
	cli, err := p.Select(in4)
	if err != nil {
		err = fmt.Errorf("api StateGetRandomnessFromTickets %v", err)
		return
	}
	return cli.StateGetRandomnessFromTickets(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateListActors(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateListActors %v", err)
		return
	}
	return cli.StateListActors(in0, in1)
}

func (p *Proxy) StateListMiners(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateListMiners %v", err)
		return
	}
	return cli.StateListMiners(in0, in1)
}

func (p *Proxy) StateLookupID(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateLookupID %v", err)
		return
	}
	return cli.StateLookupID(in0, in1, in2)
}

func (p *Proxy) StateLookupRobustAddress(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateLookupRobustAddress %v", err)
		return
	}
	return cli.StateLookupRobustAddress(in0, in1, in2)
}

func (p *Proxy) StateMarketBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MarketBalance, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMarketBalance %v", err)
		return
	}
	return cli.StateMarketBalance(in0, in1, in2)
}

func (p *Proxy) StateMarketDeals(in0 context.Context, in1 types.TipSetKey) (out0 map[string]*api1.MarketDeal, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateMarketDeals %v", err)
		return
	}
	return cli.StateMarketDeals(in0, in1)
}

func (p *Proxy) StateMarketParticipants(in0 context.Context, in1 types.TipSetKey) (out0 map[string]api1.MarketBalance, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateMarketParticipants %v", err)
		return
	}
	return cli.StateMarketParticipants(in0, in1)
}

func (p *Proxy) StateMarketStorageDeal(in0 context.Context, in1 abi.DealID, in2 types.TipSetKey) (out0 *api1.MarketDeal, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMarketStorageDeal %v", err)
		return
	}
	return cli.StateMarketStorageDeal(in0, in1, in2)
}

func (p *Proxy) StateMinerActiveSectors(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerActiveSectors %v", err)
		return
	}
	return cli.StateMinerActiveSectors(in0, in1, in2)
}

func (p *Proxy) StateMinerAllocated(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *bitfield.BitField, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerAllocated %v", err)
		return
	}
	return cli.StateMinerAllocated(in0, in1, in2)
}

func (p *Proxy) StateMinerAvailableBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerAvailableBalance %v", err)
		return
	}
	return cli.StateMinerAvailableBalance(in0, in1, in2)
}

func (p *Proxy) StateMinerDeadlines(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []api1.Deadline, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerDeadlines %v", err)
		return
	}
	return cli.StateMinerDeadlines(in0, in1, in2)
}

func (p *Proxy) StateMinerFaults(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerFaults %v", err)
		return
	}
	return cli.StateMinerFaults(in0, in1, in2)
}

func (p *Proxy) StateMinerInfo(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MinerInfo, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerInfo %v", err)
		return
	}
	return cli.StateMinerInfo(in0, in1, in2)
}

func (p *Proxy) StateMinerInitialPledgeCollateral(in0 context.Context, in1 address.Address, in2 miner.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateMinerInitialPledgeCollateral %v", err)
		return
	}
	return cli.StateMinerInitialPledgeCollateral(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerPartitions(in0 context.Context, in1 address.Address, in2 uint64, in3 types.TipSetKey) (out0 []api1.Partition, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateMinerPartitions %v", err)
		return
	}
	return cli.StateMinerPartitions(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerPower(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.MinerPower, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerPower %v", err)
		return
	}
	return cli.StateMinerPower(in0, in1, in2)
}

func (p *Proxy) StateMinerPreCommitDepositForPower(in0 context.Context, in1 address.Address, in2 miner.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateMinerPreCommitDepositForPower %v", err)
		return
	}
	return cli.StateMinerPreCommitDepositForPower(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerProvingDeadline(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *dline.Info, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerProvingDeadline %v", err)
		return
	}
	return cli.StateMinerProvingDeadline(in0, in1, in2)
}

func (p *Proxy) StateMinerRecoveries(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerRecoveries %v", err)
		return
	}
	return cli.StateMinerRecoveries(in0, in1, in2)
}

func (p *Proxy) StateMinerSectorAllocated(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 bool, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateMinerSectorAllocated %v", err)
		return
	}
	return cli.StateMinerSectorAllocated(in0, in1, in2, in3)
}

func (p *Proxy) StateMinerSectorCount(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MinerSectors, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateMinerSectorCount %v", err)
		return
	}
	return cli.StateMinerSectorCount(in0, in1, in2)
}

func (p *Proxy) StateMinerSectors(in0 context.Context, in1 address.Address, in2 *bitfield.BitField, in3 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateMinerSectors %v", err)
		return
	}
	return cli.StateMinerSectors(in0, in1, in2, in3)
}

func (p *Proxy) StateNetworkName(in0 context.Context) (out0 dtypes.NetworkName, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateNetworkName %v", err)
		return
	}
	return cli.StateNetworkName(in0)
}

func (p *Proxy) StateNetworkVersion(in0 context.Context, in1 types.TipSetKey) (out0 network.Version, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateNetworkVersion %v", err)
		return
	}
	return cli.StateNetworkVersion(in0, in1)
}

func (p *Proxy) StateReadState(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.ActorState, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateReadState %v", err)
		return
	}
	return cli.StateReadState(in0, in1, in2)
}

func (p *Proxy) StateSearchMsg(in0 context.Context, in1 types.TipSetKey, in2 cid.Cid, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateSearchMsg %v", err)
		return
	}
	return cli.StateSearchMsg(in0, in1, in2, in3, in4)
}

func (p *Proxy) StateSectorExpiration(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner1.SectorExpiration, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateSectorExpiration %v", err)
		return
	}
	return cli.StateSectorExpiration(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorGetInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorOnChainInfo, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateSectorGetInfo %v", err)
		return
	}
	return cli.StateSectorGetInfo(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorPartition(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner1.SectorLocation, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateSectorPartition %v", err)
		return
	}
	return cli.StateSectorPartition(in0, in1, in2, in3)
}

func (p *Proxy) StateSectorPreCommitInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorPreCommitOnChainInfo, err error) {
	cli, err := p.Select(in3)
	if err != nil {
		err = fmt.Errorf("api StateSectorPreCommitInfo %v", err)
		return
	}
	return cli.StateSectorPreCommitInfo(in0, in1, in2, in3)
}

func (p *Proxy) StateVMCirculatingSupplyInternal(in0 context.Context, in1 types.TipSetKey) (out0 api1.CirculatingSupply, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateVMCirculatingSupplyInternal %v", err)
		return
	}
	return cli.StateVMCirculatingSupplyInternal(in0, in1)
}

func (p *Proxy) StateVerifiedClientStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateVerifiedClientStatus %v", err)
		return
	}
	return cli.StateVerifiedClientStatus(in0, in1, in2)
}

func (p *Proxy) StateVerifiedRegistryRootKey(in0 context.Context, in1 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select(in1)
	if err != nil {
		err = fmt.Errorf("api StateVerifiedRegistryRootKey %v", err)
		return
	}
	return cli.StateVerifiedRegistryRootKey(in0, in1)
}

func (p *Proxy) StateVerifierStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select(in2)
	if err != nil {
		err = fmt.Errorf("api StateVerifierStatus %v", err)
		return
	}
	return cli.StateVerifierStatus(in0, in1, in2)
}

func (p *Proxy) StateWaitMsg(in0 context.Context, in1 cid.Cid, in2 uint64, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api StateWaitMsg %v", err)
		return
	}
	return cli.StateWaitMsg(in0, in1, in2, in3, in4)
}

func (p *Proxy) SyncState(in0 context.Context) (out0 *api1.SyncState, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api SyncState %v", err)
		return
	}
	return cli.SyncState(in0)
}

func (p *Proxy) SyncSubmitBlock(in0 context.Context, in1 *types.BlockMsg) (err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api SyncSubmitBlock %v", err)
		return
	}
	return cli.SyncSubmitBlock(in0, in1)
}

func (p *Proxy) Version(in0 context.Context) (out0 api1.APIVersion, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api Version %v", err)
		return
	}
	return cli.Version(in0)
}

func (p *Proxy) WalletBalance(in0 context.Context, in1 address.Address) (out0 big.Int, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api WalletBalance %v", err)
		return
	}
	return cli.WalletBalance(in0, in1)
}

func (p *Proxy) WalletHas(in0 context.Context, in1 address.Address) (out0 bool, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api WalletHas %v", err)
		return
	}
	return cli.WalletHas(in0, in1)
}

func (p *Proxy) WalletSign(in0 context.Context, in1 address.Address, in2 []uint8) (out0 *crypto.Signature, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api WalletSign %v", err)
		return
	}
	return cli.WalletSign(in0, in1, in2)
}

func (p *Proxy) Web3ClientVersion(in0 context.Context) (out0 string, err error) {
	cli, err := p.Select(types.EmptyTSK)
	if err != nil {
		err = fmt.Errorf("api Web3ClientVersion %v", err)
		return
	}
	return cli.Web3ClientVersion(in0)
}
