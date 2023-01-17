package root

import (
	"fmt"
	"os"

	"github.com/archieneko/archiechain/command/backup"
	"github.com/archieneko/archiechain/command/genesis"
	"github.com/archieneko/archiechain/command/helper"
	"github.com/archieneko/archiechain/command/ibft"
	"github.com/archieneko/archiechain/command/license"
	"github.com/archieneko/archiechain/command/monitor"
	"github.com/archieneko/archiechain/command/peers"
	"github.com/archieneko/archiechain/command/secrets"
	"github.com/archieneko/archiechain/command/server"
	"github.com/archieneko/archiechain/command/status"
	"github.com/archieneko/archiechain/command/txpool"
	"github.com/archieneko/archiechain/command/version"
	"github.com/archieneko/archiechain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
