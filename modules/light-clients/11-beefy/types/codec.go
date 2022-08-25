package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/ibc-go/v5/modules/core/exported"
)

// RegisterInterfaces registers the beefy concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&Misbehaviour{},
	)
}
