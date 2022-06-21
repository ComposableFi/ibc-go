package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

// CheckMisbehaviourAndUpdateState determines whether or not two conflicting
// headers at the same height would have convinced the light client.
//
// NOTE: consensusState1 is the trusted consensus state that corresponds to the TrustedHeight
// of misbehaviour.Header1
// Similarly, consensusState2 is the trusted consensus state that corresponds
// to misbehaviour.Header2
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	ctx sdk.Context,
	cdc codec.BinaryCodec,
	clientStore sdk.KVStore,
	misbehaviour exported.Misbehaviour,
) (exported.ClientState, error) {
	beefyMisbehaviour, ok := misbehaviour.(*Misbehaviour)
	if !ok {
		return nil, sdkerrors.Wrapf(clienttypes.ErrInvalidClientType, "expected type %T, got %T", misbehaviour, &Misbehaviour{})
	}

	// TODO: implement misbehaviour detection logic

	return &cs, nil
}

// verifyMisbehaviour determines whether or not two conflicting
// headers at the same height would have convinced the light client.
//
func (cs *ClientState) verifyMisbehaviour(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec, misbehaviour *Misbehaviour) error {

	return nil
}

// checkMisbehaviourHeader checks that a Header in Misbehaviour is valid misbehaviour given
// a trusted ConsensusState
func checkMisbehaviourHeader(
	clientState *ClientState, consState *ConsensusState, header *Header, currentTimestamp time.Time,
) error {

	return nil
}
