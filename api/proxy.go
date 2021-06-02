package api

import (
	"context"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

var _ Proxy = api.FullNode(nil)

// Proxy is a subset of api.FullNode, which provides the required methods for a chain service
type Proxy interface {
	// ChainHead returns the current head of the chain.
	ChainHead(context.Context) (*types.TipSet, error)

	// ChainGetRandomnessFromTickets is used to sample the chain for randomness.
	ChainGetRandomnessFromTickets(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)

	// ChainGetRandomnessFromBeacon is used to sample the beacon for randomness.
	ChainGetRandomnessFromBeacon(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)

	// ChainGetBlock returns the block specified by the given CID.
	ChainGetBlock(context.Context, cid.Cid) (*types.BlockHeader, error)
	// ChainGetTipSet returns the tipset specified by the given TipSetKey.
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)

	// ChainGetBlockMessages returns messages stored in the specified block.
	ChainGetBlockMessages(ctx context.Context, blockCid cid.Cid) (*api.BlockMessages, error)

	// ChainGetParentReceipts returns receipts for messages in parent tipset of
	// the specified block.
	ChainGetParentReceipts(ctx context.Context, blockCid cid.Cid) ([]*types.MessageReceipt, error)

	// ChainGetParentMessages returns messages stored in parent tipset of the
	// specified block.
	ChainGetParentMessages(ctx context.Context, blockCid cid.Cid) ([]api.Message, error)

	// ChainGetTipSetByHeight looks back for a tipset at the specified epoch.
	// If there are no blocks at the specified epoch, a tipset at an earlier epoch
	// will be returned.
	ChainGetTipSetByHeight(context.Context, abi.ChainEpoch, types.TipSetKey) (*types.TipSet, error)

	// ChainGetGenesis returns the genesis tipset.
	ChainGetGenesis(context.Context) (*types.TipSet, error)

	// ChainTipSetWeight computes weight for the specified tipset.
	ChainTipSetWeight(context.Context, types.TipSetKey) (types.BigInt, error)

	// ChainGetMessage reads a message referenced by the specified CID from the
	// chain blockstore.
	ChainGetMessage(context.Context, cid.Cid) (*types.Message, error)

	// MethodGroup: Beacon
	// The Beacon method group contains methods for interacting with the random beacon (DRAND)

	// BeaconGetEntry returns the beacon entry for the given filecoin epoch. If
	// the entry has not yet been produced, the call will block until the entry
	// becomes available
	BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error)
}
