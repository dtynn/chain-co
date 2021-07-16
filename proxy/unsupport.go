package proxy

import (
	"context"
	"github.com/dtynn/chain-co/api"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-multistore"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/dline"
	network1 "github.com/filecoin-project/go-state-types/network"
	api1 "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/types"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/markets/loggers"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	miner1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	"github.com/filecoin-project/specs-actors/actors/builtin/paych"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

var _ UnSupportAPI = (*UnSupport)(nil)

type UnSupportAPI interface {
	api.UnSupport
}

type UnSupport struct {
	Select func() (UnSupportAPI, error)
}

// impl api.UnSupport
func (p *UnSupport) AuthNew(in0 context.Context, in1 []auth.Permission) (out0 []uint8, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.AuthNew(in0, in1)
}

func (p *UnSupport) AuthVerify(in0 context.Context, in1 string) (out0 []auth.Permission, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.AuthVerify(in0, in1)
}

func (p *UnSupport) ChainDeleteObj(in0 context.Context, in1 cid.Cid) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainDeleteObj(in0, in1)
}

func (p *UnSupport) ChainExport(in0 context.Context, in1 abi.ChainEpoch, in2 bool, in3 types.TipSetKey) (out0 <-chan []uint8, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainExport(in0, in1, in2, in3)
}

func (p *UnSupport) ChainGetNode(in0 context.Context, in1 string) (out0 *api1.IpldObject, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetNode(in0, in1)
}

func (p *UnSupport) ChainGetPath(in0 context.Context, in1 types.TipSetKey, in2 types.TipSetKey) (out0 []*api1.HeadChange, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainGetPath(in0, in1, in2)
}

func (p *UnSupport) ChainHasObj(in0 context.Context, in1 cid.Cid) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainHasObj(in0, in1)
}

func (p *UnSupport) ChainReadObj(in0 context.Context, in1 cid.Cid) (out0 []uint8, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainReadObj(in0, in1)
}

func (p *UnSupport) ChainSetHead(in0 context.Context, in1 types.TipSetKey) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainSetHead(in0, in1)
}

func (p *UnSupport) ChainStatObj(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 api1.ObjStat, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ChainStatObj(in0, in1, in2)
}

func (p *UnSupport) ClientCalcCommP(in0 context.Context, in1 string) (out0 *api1.CommPRet, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientCalcCommP(in0, in1)
}

func (p *UnSupport) ClientCancelDataTransfer(in0 context.Context, in1 datatransfer.TransferID, in2 peer.ID, in3 bool) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientCancelDataTransfer(in0, in1, in2, in3)
}

func (p *UnSupport) ClientCancelRetrievalDeal(in0 context.Context, in1 retrievalmarket.DealID) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientCancelRetrievalDeal(in0, in1)
}

func (p *UnSupport) ClientDataTransferUpdates(in0 context.Context) (out0 <-chan api1.DataTransferChannel, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientDataTransferUpdates(in0)
}

func (p *UnSupport) ClientDealPieceCID(in0 context.Context, in1 cid.Cid) (out0 api1.DataCIDSize, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientDealPieceCID(in0, in1)
}

func (p *UnSupport) ClientDealSize(in0 context.Context, in1 cid.Cid) (out0 api1.DataSize, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientDealSize(in0, in1)
}

func (p *UnSupport) ClientFindData(in0 context.Context, in1 cid.Cid, in2 *cid.Cid) (out0 []api1.QueryOffer, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientFindData(in0, in1, in2)
}

func (p *UnSupport) ClientGenCar(in0 context.Context, in1 api1.FileRef, in2 string) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientGenCar(in0, in1, in2)
}

func (p *UnSupport) ClientGetDealInfo(in0 context.Context, in1 cid.Cid) (out0 *api1.DealInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientGetDealInfo(in0, in1)
}

func (p *UnSupport) ClientGetDealStatus(in0 context.Context, in1 uint64) (out0 string, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientGetDealStatus(in0, in1)
}

func (p *UnSupport) ClientGetDealUpdates(in0 context.Context) (out0 <-chan api1.DealInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientGetDealUpdates(in0)
}

func (p *UnSupport) ClientHasLocal(in0 context.Context, in1 cid.Cid) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientHasLocal(in0, in1)
}

func (p *UnSupport) ClientImport(in0 context.Context, in1 api1.FileRef) (out0 *api1.ImportRes, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientImport(in0, in1)
}

func (p *UnSupport) ClientListDataTransfers(in0 context.Context) (out0 []api1.DataTransferChannel, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientListDataTransfers(in0)
}

func (p *UnSupport) ClientListDeals(in0 context.Context) (out0 []api1.DealInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientListDeals(in0)
}

func (p *UnSupport) ClientListImports(in0 context.Context) (out0 []api1.Import, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientListImports(in0)
}

func (p *UnSupport) ClientMinerQueryOffer(in0 context.Context, in1 address.Address, in2 cid.Cid, in3 *cid.Cid) (out0 api1.QueryOffer, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientMinerQueryOffer(in0, in1, in2, in3)
}

func (p *UnSupport) ClientQueryAsk(in0 context.Context, in1 peer.ID, in2 address.Address) (out0 *storagemarket.StorageAsk, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientQueryAsk(in0, in1, in2)
}

func (p *UnSupport) ClientRemoveImport(in0 context.Context, in1 multistore.StoreID) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientRemoveImport(in0, in1)
}

func (p *UnSupport) ClientRestartDataTransfer(in0 context.Context, in1 datatransfer.TransferID, in2 peer.ID, in3 bool) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientRestartDataTransfer(in0, in1, in2, in3)
}

func (p *UnSupport) ClientRetrieve(in0 context.Context, in1 api1.RetrievalOrder, in2 *api1.FileRef) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientRetrieve(in0, in1, in2)
}

func (p *UnSupport) ClientRetrieveTryRestartInsufficientFunds(in0 context.Context, in1 address.Address) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientRetrieveTryRestartInsufficientFunds(in0, in1)
}

func (p *UnSupport) ClientRetrieveWithEvents(in0 context.Context, in1 api1.RetrievalOrder, in2 *api1.FileRef) (out0 <-chan marketevents.RetrievalEvent, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientRetrieveWithEvents(in0, in1, in2)
}

func (p *UnSupport) ClientStartDeal(in0 context.Context, in1 *api1.StartDealParams) (out0 *cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ClientStartDeal(in0, in1)
}

func (p *UnSupport) Closing(in0 context.Context) (out0 <-chan struct{}, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.Closing(in0)
}

func (p *UnSupport) CreateBackup(in0 context.Context, in1 string) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.CreateBackup(in0, in1)
}

func (p *UnSupport) Discover(in0 context.Context) (out0 apitypes.OpenRPCDocument, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.Discover(in0)
}

func (p *UnSupport) GasEstimateFeeCap(in0 context.Context, in1 *types.Message, in2 int64, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.GasEstimateFeeCap(in0, in1, in2, in3)
}

func (p *UnSupport) GasEstimateGasLimit(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 int64, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.GasEstimateGasLimit(in0, in1, in2)
}

func (p *UnSupport) GasEstimateGasPremium(in0 context.Context, in1 uint64, in2 address.Address, in3 int64, in4 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.GasEstimateGasPremium(in0, in1, in2, in3, in4)
}

func (p *UnSupport) GasEstimateMessageGas(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec, in3 types.TipSetKey) (out0 *types.Message, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.GasEstimateMessageGas(in0, in1, in2, in3)
}

func (p *UnSupport) ID(in0 context.Context) (out0 peer.ID, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.ID(in0)
}

func (p *UnSupport) LogList(in0 context.Context) (out0 []string, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.LogList(in0)
}

func (p *UnSupport) LogSetLevel(in0 context.Context, in1 string, in2 string) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.LogSetLevel(in0, in1, in2)
}

func (p *UnSupport) MarketAddBalance(in0 context.Context, in1 address.Address, in2 address.Address, in3 big.Int) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MarketAddBalance(in0, in1, in2, in3)
}

func (p *UnSupport) MarketGetReserved(in0 context.Context, in1 address.Address) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MarketGetReserved(in0, in1)
}

func (p *UnSupport) MarketReleaseFunds(in0 context.Context, in1 address.Address, in2 big.Int) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MarketReleaseFunds(in0, in1, in2)
}

func (p *UnSupport) MarketReserveFunds(in0 context.Context, in1 address.Address, in2 address.Address, in3 big.Int) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MarketReserveFunds(in0, in1, in2, in3)
}

func (p *UnSupport) MarketWithdraw(in0 context.Context, in1 address.Address, in2 address.Address, in3 big.Int) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MarketWithdraw(in0, in1, in2, in3)
}

func (p *UnSupport) MinerCreateBlock(in0 context.Context, in1 *api1.BlockTemplate) (out0 *types.BlockMsg, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MinerCreateBlock(in0, in1)
}

func (p *UnSupport) MinerGetBaseInfo(in0 context.Context, in1 address.Address, in2 abi.ChainEpoch, in3 types.TipSetKey) (out0 *api1.MiningBaseInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MinerGetBaseInfo(in0, in1, in2, in3)
}

func (p *UnSupport) MpoolBatchPush(in0 context.Context, in1 []*types.SignedMessage) (out0 []cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolBatchPush(in0, in1)
}

func (p *UnSupport) MpoolBatchPushMessage(in0 context.Context, in1 []*types.Message, in2 *api1.MessageSendSpec) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolBatchPushMessage(in0, in1, in2)
}

func (p *UnSupport) MpoolBatchPushUntrusted(in0 context.Context, in1 []*types.SignedMessage) (out0 []cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolBatchPushUntrusted(in0, in1)
}

func (p *UnSupport) MpoolClear(in0 context.Context, in1 bool) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolClear(in0, in1)
}

func (p *UnSupport) MpoolGetConfig(in0 context.Context) (out0 *types.MpoolConfig, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolGetConfig(in0)
}

func (p *UnSupport) MpoolGetNonce(in0 context.Context, in1 address.Address) (out0 uint64, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolGetNonce(in0, in1)
}

func (p *UnSupport) MpoolPending(in0 context.Context, in1 types.TipSetKey) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolPending(in0, in1)
}

func (p *UnSupport) MpoolPush(in0 context.Context, in1 *types.SignedMessage) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolPush(in0, in1)
}

func (p *UnSupport) MpoolPushMessage(in0 context.Context, in1 *types.Message, in2 *api1.MessageSendSpec) (out0 *types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolPushMessage(in0, in1, in2)
}

func (p *UnSupport) MpoolPushUntrusted(in0 context.Context, in1 *types.SignedMessage) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolPushUntrusted(in0, in1)
}

func (p *UnSupport) MpoolSelect(in0 context.Context, in1 types.TipSetKey, in2 float64) (out0 []*types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolSelect(in0, in1, in2)
}

func (p *UnSupport) MpoolSetConfig(in0 context.Context, in1 *types.MpoolConfig) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolSetConfig(in0, in1)
}

func (p *UnSupport) MpoolSub(in0 context.Context) (out0 <-chan api1.MpoolUpdate, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MpoolSub(in0)
}

func (p *UnSupport) MsigAddApprove(in0 context.Context, in1 address.Address, in2 address.Address, in3 uint64, in4 address.Address, in5 address.Address, in6 bool) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigAddApprove(in0, in1, in2, in3, in4, in5, in6)
}

func (p *UnSupport) MsigAddCancel(in0 context.Context, in1 address.Address, in2 address.Address, in3 uint64, in4 address.Address, in5 bool) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigAddCancel(in0, in1, in2, in3, in4, in5)
}

func (p *UnSupport) MsigAddPropose(in0 context.Context, in1 address.Address, in2 address.Address, in3 address.Address, in4 bool) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigAddPropose(in0, in1, in2, in3, in4)
}

func (p *UnSupport) MsigApprove(in0 context.Context, in1 address.Address, in2 uint64, in3 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigApprove(in0, in1, in2, in3)
}

func (p *UnSupport) MsigApproveTxnHash(in0 context.Context, in1 address.Address, in2 uint64, in3 address.Address, in4 address.Address, in5 big.Int, in6 address.Address, in7 uint64, in8 []uint8) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigApproveTxnHash(in0, in1, in2, in3, in4, in5, in6, in7, in8)
}

func (p *UnSupport) MsigCancel(in0 context.Context, in1 address.Address, in2 uint64, in3 address.Address, in4 big.Int, in5 address.Address, in6 uint64, in7 []uint8) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigCancel(in0, in1, in2, in3, in4, in5, in6, in7)
}

func (p *UnSupport) MsigCreate(in0 context.Context, in1 uint64, in2 []address.Address, in3 abi.ChainEpoch, in4 big.Int, in5 address.Address, in6 big.Int) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigCreate(in0, in1, in2, in3, in4, in5, in6)
}

func (p *UnSupport) MsigGetAvailableBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigGetAvailableBalance(in0, in1, in2)
}

func (p *UnSupport) MsigGetPending(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []*api1.MsigTransaction, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigGetPending(in0, in1, in2)
}

func (p *UnSupport) MsigGetVested(in0 context.Context, in1 address.Address, in2 types.TipSetKey, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigGetVested(in0, in1, in2, in3)
}

func (p *UnSupport) MsigGetVestingSchedule(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MsigVesting, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigGetVestingSchedule(in0, in1, in2)
}

func (p *UnSupport) MsigPropose(in0 context.Context, in1 address.Address, in2 address.Address, in3 big.Int, in4 address.Address, in5 uint64, in6 []uint8) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigPropose(in0, in1, in2, in3, in4, in5, in6)
}

func (p *UnSupport) MsigRemoveSigner(in0 context.Context, in1 address.Address, in2 address.Address, in3 address.Address, in4 bool) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigRemoveSigner(in0, in1, in2, in3, in4)
}

func (p *UnSupport) MsigSwapApprove(in0 context.Context, in1 address.Address, in2 address.Address, in3 uint64, in4 address.Address, in5 address.Address, in6 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigSwapApprove(in0, in1, in2, in3, in4, in5, in6)
}

func (p *UnSupport) MsigSwapCancel(in0 context.Context, in1 address.Address, in2 address.Address, in3 uint64, in4 address.Address, in5 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigSwapCancel(in0, in1, in2, in3, in4, in5)
}

func (p *UnSupport) MsigSwapPropose(in0 context.Context, in1 address.Address, in2 address.Address, in3 address.Address, in4 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.MsigSwapPropose(in0, in1, in2, in3, in4)
}

func (p *UnSupport) NetAddrsListen(in0 context.Context) (out0 peer.AddrInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetAddrsListen(in0)
}

func (p *UnSupport) NetAgentVersion(in0 context.Context, in1 peer.ID) (out0 string, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetAgentVersion(in0, in1)
}

func (p *UnSupport) NetAutoNatStatus(in0 context.Context) (out0 api1.NatInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetAutoNatStatus(in0)
}

func (p *UnSupport) NetBandwidthStats(in0 context.Context) (out0 metrics.Stats, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBandwidthStats(in0)
}

func (p *UnSupport) NetBandwidthStatsByPeer(in0 context.Context) (out0 map[string]metrics.Stats, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBandwidthStatsByPeer(in0)
}

func (p *UnSupport) NetBandwidthStatsByProtocol(in0 context.Context) (out0 map[protocol.ID]metrics.Stats, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBandwidthStatsByProtocol(in0)
}

func (p *UnSupport) NetBlockAdd(in0 context.Context, in1 api1.NetBlockList) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBlockAdd(in0, in1)
}

func (p *UnSupport) NetBlockList(in0 context.Context) (out0 api1.NetBlockList, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBlockList(in0)
}

func (p *UnSupport) NetBlockRemove(in0 context.Context, in1 api1.NetBlockList) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetBlockRemove(in0, in1)
}

func (p *UnSupport) NetConnect(in0 context.Context, in1 peer.AddrInfo) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetConnect(in0, in1)
}

func (p *UnSupport) NetConnectedness(in0 context.Context, in1 peer.ID) (out0 network.Connectedness, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetConnectedness(in0, in1)
}

func (p *UnSupport) NetDisconnect(in0 context.Context, in1 peer.ID) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetDisconnect(in0, in1)
}

func (p *UnSupport) NetFindPeer(in0 context.Context, in1 peer.ID) (out0 peer.AddrInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetFindPeer(in0, in1)
}

func (p *UnSupport) NetPeerInfo(in0 context.Context, in1 peer.ID) (out0 *api1.ExtendedPeerInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetPeerInfo(in0, in1)
}

func (p *UnSupport) NetPeers(in0 context.Context) (out0 []peer.AddrInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetPeers(in0)
}

func (p *UnSupport) NetPubsubScores(in0 context.Context) (out0 []api1.PubsubScore, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.NetPubsubScores(in0)
}

func (p *UnSupport) PaychAllocateLane(in0 context.Context, in1 address.Address) (out0 uint64, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychAllocateLane(in0, in1)
}

func (p *UnSupport) PaychAvailableFunds(in0 context.Context, in1 address.Address) (out0 *api1.ChannelAvailableFunds, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychAvailableFunds(in0, in1)
}

func (p *UnSupport) PaychAvailableFundsByFromTo(in0 context.Context, in1 address.Address, in2 address.Address) (out0 *api1.ChannelAvailableFunds, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychAvailableFundsByFromTo(in0, in1, in2)
}

func (p *UnSupport) PaychCollect(in0 context.Context, in1 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychCollect(in0, in1)
}

func (p *UnSupport) PaychGet(in0 context.Context, in1 address.Address, in2 address.Address, in3 big.Int) (out0 *api1.ChannelInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychGet(in0, in1, in2, in3)
}

func (p *UnSupport) PaychGetWaitReady(in0 context.Context, in1 cid.Cid) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychGetWaitReady(in0, in1)
}

func (p *UnSupport) PaychList(in0 context.Context) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychList(in0)
}

func (p *UnSupport) PaychNewPayment(in0 context.Context, in1 address.Address, in2 address.Address, in3 []api1.VoucherSpec) (out0 *api1.PaymentInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychNewPayment(in0, in1, in2, in3)
}

func (p *UnSupport) PaychSettle(in0 context.Context, in1 address.Address) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychSettle(in0, in1)
}

func (p *UnSupport) PaychStatus(in0 context.Context, in1 address.Address) (out0 *api1.PaychStatus, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychStatus(in0, in1)
}

func (p *UnSupport) PaychVoucherAdd(in0 context.Context, in1 address.Address, in2 *paych.SignedVoucher, in3 []uint8, in4 big.Int) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherAdd(in0, in1, in2, in3, in4)
}

func (p *UnSupport) PaychVoucherCheckSpendable(in0 context.Context, in1 address.Address, in2 *paych.SignedVoucher, in3 []uint8, in4 []uint8) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherCheckSpendable(in0, in1, in2, in3, in4)
}

func (p *UnSupport) PaychVoucherCheckValid(in0 context.Context, in1 address.Address, in2 *paych.SignedVoucher) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherCheckValid(in0, in1, in2)
}

func (p *UnSupport) PaychVoucherCreate(in0 context.Context, in1 address.Address, in2 big.Int, in3 uint64) (out0 *api1.VoucherCreateResult, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherCreate(in0, in1, in2, in3)
}

func (p *UnSupport) PaychVoucherList(in0 context.Context, in1 address.Address) (out0 []*paych.SignedVoucher, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherList(in0, in1)
}

func (p *UnSupport) PaychVoucherSubmit(in0 context.Context, in1 address.Address, in2 *paych.SignedVoucher, in3 []uint8, in4 []uint8) (out0 cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.PaychVoucherSubmit(in0, in1, in2, in3, in4)
}

func (p *UnSupport) Session(in0 context.Context) (out0 uuid.UUID, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.Session(in0)
}

func (p *UnSupport) Shutdown(in0 context.Context) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.Shutdown(in0)
}

func (p *UnSupport) StateAccountKey(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateAccountKey(in0, in1, in2)
}

func (p *UnSupport) StateAllMinerFaults(in0 context.Context, in1 abi.ChainEpoch, in2 types.TipSetKey) (out0 []*api1.Fault, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateAllMinerFaults(in0, in1, in2)
}

func (p *UnSupport) StateCall(in0 context.Context, in1 *types.Message, in2 types.TipSetKey) (out0 *api1.InvocResult, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateCall(in0, in1, in2)
}

func (p *UnSupport) StateChangedActors(in0 context.Context, in1 cid.Cid, in2 cid.Cid) (out0 map[string]types.Actor, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateChangedActors(in0, in1, in2)
}

func (p *UnSupport) StateCirculatingSupply(in0 context.Context, in1 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateCirculatingSupply(in0, in1)
}

func (p *UnSupport) StateCompute(in0 context.Context, in1 abi.ChainEpoch, in2 []*types.Message, in3 types.TipSetKey) (out0 *api1.ComputeStateOutput, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateCompute(in0, in1, in2, in3)
}

func (p *UnSupport) StateDealProviderCollateralBounds(in0 context.Context, in1 abi.PaddedPieceSize, in2 bool, in3 types.TipSetKey) (out0 api1.DealCollateralBounds, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateDealProviderCollateralBounds(in0, in1, in2, in3)
}

func (p *UnSupport) StateDecodeParams(in0 context.Context, in1 address.Address, in2 abi.MethodNum, in3 []uint8, in4 types.TipSetKey) (out0 interface{}, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateDecodeParams(in0, in1, in2, in3, in4)
}

func (p *UnSupport) StateGetActor(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *types.Actor, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateGetActor(in0, in1, in2)
}

func (p *UnSupport) StateGetReceipt(in0 context.Context, in1 cid.Cid, in2 types.TipSetKey) (out0 *types.MessageReceipt, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateGetReceipt(in0, in1, in2)
}

func (p *UnSupport) StateListActors(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateListActors(in0, in1)
}

func (p *UnSupport) StateListMessages(in0 context.Context, in1 *api1.MessageMatch, in2 types.TipSetKey, in3 abi.ChainEpoch) (out0 []cid.Cid, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateListMessages(in0, in1, in2, in3)
}

func (p *UnSupport) StateListMiners(in0 context.Context, in1 types.TipSetKey) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateListMiners(in0, in1)
}

func (p *UnSupport) StateLookupID(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateLookupID(in0, in1, in2)
}

func (p *UnSupport) StateMarketBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MarketBalance, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMarketBalance(in0, in1, in2)
}

func (p *UnSupport) StateMarketDeals(in0 context.Context, in1 types.TipSetKey) (out0 map[string]api1.MarketDeal, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMarketDeals(in0, in1)
}

func (p *UnSupport) StateMarketParticipants(in0 context.Context, in1 types.TipSetKey) (out0 map[string]api1.MarketBalance, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMarketParticipants(in0, in1)
}

func (p *UnSupport) StateMarketStorageDeal(in0 context.Context, in1 abi.DealID, in2 types.TipSetKey) (out0 *api1.MarketDeal, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMarketStorageDeal(in0, in1, in2)
}

func (p *UnSupport) StateMinerActiveSectors(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerActiveSectors(in0, in1, in2)
}

func (p *UnSupport) StateMinerAvailableBalance(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerAvailableBalance(in0, in1, in2)
}

func (p *UnSupport) StateMinerDeadlines(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 []api1.Deadline, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerDeadlines(in0, in1, in2)
}

func (p *UnSupport) StateMinerFaults(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerFaults(in0, in1, in2)
}

func (p *UnSupport) StateMinerInfo(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 miner.MinerInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerInfo(in0, in1, in2)
}

func (p *UnSupport) StateMinerInitialPledgeCollateral(in0 context.Context, in1 address.Address, in2 miner1.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerInitialPledgeCollateral(in0, in1, in2, in3)
}

func (p *UnSupport) StateMinerPartitions(in0 context.Context, in1 address.Address, in2 uint64, in3 types.TipSetKey) (out0 []api1.Partition, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerPartitions(in0, in1, in2, in3)
}

func (p *UnSupport) StateMinerPower(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.MinerPower, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerPower(in0, in1, in2)
}

func (p *UnSupport) StateMinerPreCommitDepositForPower(in0 context.Context, in1 address.Address, in2 miner1.SectorPreCommitInfo, in3 types.TipSetKey) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerPreCommitDepositForPower(in0, in1, in2, in3)
}

func (p *UnSupport) StateMinerProvingDeadline(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *dline.Info, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerProvingDeadline(in0, in1, in2)
}

func (p *UnSupport) StateMinerRecoveries(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 bitfield.BitField, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerRecoveries(in0, in1, in2)
}

func (p *UnSupport) StateMinerSectorAllocated(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerSectorAllocated(in0, in1, in2, in3)
}

func (p *UnSupport) StateMinerSectorCount(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 api1.MinerSectors, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerSectorCount(in0, in1, in2)
}

func (p *UnSupport) StateMinerSectors(in0 context.Context, in1 address.Address, in2 *bitfield.BitField, in3 types.TipSetKey) (out0 []*miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateMinerSectors(in0, in1, in2, in3)
}

func (p *UnSupport) StateNetworkName(in0 context.Context) (out0 dtypes.NetworkName, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateNetworkName(in0)
}

func (p *UnSupport) StateNetworkVersion(in0 context.Context, in1 types.TipSetKey) (out0 network1.Version, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateNetworkVersion(in0, in1)
}

func (p *UnSupport) StateReadState(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *api1.ActorState, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateReadState(in0, in1, in2)
}

func (p *UnSupport) StateReplay(in0 context.Context, in1 types.TipSetKey, in2 cid.Cid) (out0 *api1.InvocResult, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateReplay(in0, in1, in2)
}

func (p *UnSupport) StateSearchMsg(in0 context.Context, in1 types.TipSetKey, in2 cid.Cid, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSearchMsg(in0, in1, in2, in3, in4)
}

func (p *UnSupport) StateSearchMsgLimited(in0 context.Context, in1 cid.Cid, in2 abi.ChainEpoch) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSearchMsgLimited(in0, in1, in2)
}

func (p *UnSupport) StateSectorExpiration(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorExpiration, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSectorExpiration(in0, in1, in2, in3)
}

func (p *UnSupport) StateSectorGetInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSectorGetInfo(in0, in1, in2, in3)
}

func (p *UnSupport) StateSectorPartition(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 *miner.SectorLocation, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSectorPartition(in0, in1, in2, in3)
}

func (p *UnSupport) StateSectorPreCommitInfo(in0 context.Context, in1 address.Address, in2 abi.SectorNumber, in3 types.TipSetKey) (out0 miner.SectorPreCommitOnChainInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateSectorPreCommitInfo(in0, in1, in2, in3)
}

func (p *UnSupport) StateVMCirculatingSupplyInternal(in0 context.Context, in1 types.TipSetKey) (out0 api1.CirculatingSupply, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateVMCirculatingSupplyInternal(in0, in1)
}

func (p *UnSupport) StateVerifiedClientStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateVerifiedClientStatus(in0, in1, in2)
}

func (p *UnSupport) StateVerifiedRegistryRootKey(in0 context.Context, in1 types.TipSetKey) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateVerifiedRegistryRootKey(in0, in1)
}

func (p *UnSupport) StateVerifierStatus(in0 context.Context, in1 address.Address, in2 types.TipSetKey) (out0 *big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateVerifierStatus(in0, in1, in2)
}

func (p *UnSupport) StateWaitMsg(in0 context.Context, in1 cid.Cid, in2 uint64, in3 abi.ChainEpoch, in4 bool) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateWaitMsg(in0, in1, in2, in3, in4)
}

func (p *UnSupport) StateWaitMsgLimited(in0 context.Context, in1 cid.Cid, in2 uint64, in3 abi.ChainEpoch) (out0 *api1.MsgLookup, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.StateWaitMsgLimited(in0, in1, in2, in3)
}

func (p *UnSupport) SyncCheckBad(in0 context.Context, in1 cid.Cid) (out0 string, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncCheckBad(in0, in1)
}

func (p *UnSupport) SyncCheckpoint(in0 context.Context, in1 types.TipSetKey) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncCheckpoint(in0, in1)
}

func (p *UnSupport) SyncIncomingBlocks(in0 context.Context) (out0 <-chan *types.BlockHeader, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncIncomingBlocks(in0)
}

func (p *UnSupport) SyncMarkBad(in0 context.Context, in1 cid.Cid) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncMarkBad(in0, in1)
}

func (p *UnSupport) SyncState(in0 context.Context) (out0 *api1.SyncState, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncState(in0)
}

func (p *UnSupport) SyncSubmitBlock(in0 context.Context, in1 *types.BlockMsg) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncSubmitBlock(in0, in1)
}

func (p *UnSupport) SyncUnmarkAllBad(in0 context.Context) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncUnmarkAllBad(in0)
}

func (p *UnSupport) SyncUnmarkBad(in0 context.Context, in1 cid.Cid) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncUnmarkBad(in0, in1)
}

func (p *UnSupport) SyncValidateTipset(in0 context.Context, in1 types.TipSetKey) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.SyncValidateTipset(in0, in1)
}

func (p *UnSupport) Version(in0 context.Context) (out0 api1.APIVersion, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.Version(in0)
}

func (p *UnSupport) WalletBalance(in0 context.Context, in1 address.Address) (out0 big.Int, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletBalance(in0, in1)
}

func (p *UnSupport) WalletDefaultAddress(in0 context.Context) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletDefaultAddress(in0)
}

func (p *UnSupport) WalletDelete(in0 context.Context, in1 address.Address) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletDelete(in0, in1)
}

func (p *UnSupport) WalletExport(in0 context.Context, in1 address.Address) (out0 *types.KeyInfo, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletExport(in0, in1)
}

func (p *UnSupport) WalletHas(in0 context.Context, in1 address.Address) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletHas(in0, in1)
}

func (p *UnSupport) WalletImport(in0 context.Context, in1 *types.KeyInfo) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletImport(in0, in1)
}

func (p *UnSupport) WalletList(in0 context.Context) (out0 []address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletList(in0)
}

func (p *UnSupport) WalletNew(in0 context.Context, in1 types.KeyType) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletNew(in0, in1)
}

func (p *UnSupport) WalletSetDefault(in0 context.Context, in1 address.Address) (err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletSetDefault(in0, in1)
}

func (p *UnSupport) WalletSign(in0 context.Context, in1 address.Address, in2 []uint8) (out0 *crypto.Signature, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletSign(in0, in1, in2)
}

func (p *UnSupport) WalletSignMessage(in0 context.Context, in1 address.Address, in2 *types.Message) (out0 *types.SignedMessage, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletSignMessage(in0, in1, in2)
}

func (p *UnSupport) WalletValidateAddress(in0 context.Context, in1 string) (out0 address.Address, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletValidateAddress(in0, in1)
}

func (p *UnSupport) WalletVerify(in0 context.Context, in1 address.Address, in2 []uint8, in3 *crypto.Signature) (out0 bool, err error) {
	cli, err := p.Select()
	if err != nil {
		return
	}
	return cli.WalletVerify(in0, in1, in2, in3)
}
