package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	ibcclient "github.com/cosmos/ibc-go/v5/modules/core/02-client"
	connection "github.com/cosmos/ibc-go/v5/modules/core/03-connection"
	channel "github.com/cosmos/ibc-go/v5/modules/core/04-channel"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	wasmcli "github.com/cosmos/ibc-go/v5/modules/light-clients/10-wasm/cli"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	ibcTxCmd := &cobra.Command{
		Use:                        host.ModuleName,
		Short:                      "IBC transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ibcTxCmd.AddCommand(
		ibcclient.GetTxCmd(),
		channel.GetTxCmd(),
		wasmcli.NewTxCmd(),
	)

	return ibcTxCmd
}

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group ibc queries under a subcommand
	ibcQueryCmd := &cobra.Command{
		Use:                        host.ModuleName,
		Short:                      "Querying commands for the IBC module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ibcQueryCmd.AddCommand(
		ibcclient.GetQueryCmd(),
		connection.GetQueryCmd(),
		channel.GetQueryCmd(),
		wasmcli.GetQueryCmd(),
	)

	return ibcQueryCmd
}
