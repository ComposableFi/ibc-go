package solomachine_test

import (
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v3/modules/light-clients/06-solomachine"
	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

func (suite *SoloMachineTestSuite) TestCheckSubstituteAndUpdateState() {
	var (
		subjectClientState    *solomachine.ClientState
		substituteClientState exported.ClientState
	)

	// test singlesig and multisig public keys
	for _, sm := range []*ibctesting.Solomachine{suite.solomachine, suite.solomachineMulti} {

		testCases := []struct {
			name     string
			malleate func()
			expPass  bool
		}{
			{
				"valid substitute", func() {
					subjectClientState.AllowUpdateAfterProposal = true
				}, true,
			},
			{
				"subject not allowed to be updated", func() {
					subjectClientState.AllowUpdateAfterProposal = false
				}, false,
			},
			{
				"substitute is not the solo machine", func() {
					substituteClientState = &ibctmtypes.ClientState{}
				}, false,
			},
			{
				"subject public key is nil", func() {
					subjectClientState.ConsensusState.PublicKey = nil
				}, false,
			},

			{
				"substitute public key is nil", func() {
					substituteClientState.(*solomachine.ClientState).ConsensusState.PublicKey = nil
				}, false,
			},
			{
				"subject and substitute use the same public key", func() {
					substituteClientState.(*solomachine.ClientState).ConsensusState.PublicKey = subjectClientState.ConsensusState.PublicKey
				}, false,
			},
		}

		for _, tc := range testCases {
			tc := tc

			suite.Run(tc.name, func() {
				suite.SetupTest()

				subjectClientState = sm.ClientState()
				subjectClientState.AllowUpdateAfterProposal = true
				substitute := ibctesting.NewSolomachine(suite.T(), suite.chainA.Codec, "substitute", "testing", 5)
				substituteClientState = substitute.ClientState()

				tc.malleate()

				subjectClientStore := suite.chainA.App.GetIBCKeeper().ClientKeeper.ClientStore(suite.chainA.GetContext(), sm.ClientID)
				substituteClientStore := suite.chainA.App.GetIBCKeeper().ClientKeeper.ClientStore(suite.chainA.GetContext(), substitute.ClientID)

				updatedClient, err := subjectClientState.CheckSubstituteAndUpdateState(suite.chainA.GetContext(), suite.chainA.App.AppCodec(), subjectClientStore, substituteClientStore, substituteClientState)

				if tc.expPass {
					suite.Require().NoError(err)

					suite.Require().Equal(substituteClientState.(*solomachine.ClientState).ConsensusState, updatedClient.(*solomachine.ClientState).ConsensusState)
					suite.Require().Equal(substituteClientState.(*solomachine.ClientState).Sequence, updatedClient.(*solomachine.ClientState).Sequence)
					suite.Require().Equal(false, updatedClient.(*solomachine.ClientState).IsFrozen)

					// ensure updated client state is set in store
					bz := subjectClientStore.Get(host.ClientStateKey())
					suite.Require().Equal(clienttypes.MustMarshalClientState(suite.chainA.Codec, updatedClient), bz)
				} else {
					suite.Require().Error(err)
					suite.Require().Nil(updatedClient)
				}
			})
		}
	}
}
