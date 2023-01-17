package ibft

import (
	"github.com/archieneko/archiechain/command/helper"
	"github.com/archieneko/archiechain/command/ibft/candidates"
	"github.com/archieneko/archiechain/command/ibft/propose"
	"github.com/archieneko/archiechain/command/ibft/quorum"
	"github.com/archieneko/archiechain/command/ibft/snapshot"
	"github.com/archieneko/archiechain/command/ibft/status"
	_switch "github.com/archieneko/archiechain/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
