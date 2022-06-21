package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

func (cs ClientState) VerifyUpgradeAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	clientStore sdk.KVStore,
	upgradedClient exported.ClientState,
	upgradedConsState exported.ConsensusState,
	proofUpgradeClient, proofUpgradeConsState []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	panic("implement me")
}

func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	subjectClientStore,
	substituteClientStore sdk.KVStore,
	substituteClient exported.ClientState,
) (exported.ClientState, error) {
	panic("implement me")
}
