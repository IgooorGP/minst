package cmd

import (
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/spf13/cobra"
)

// NewMinstCmd creates the `minst` command and its nested subcommands.
func NewMinstCmd() *cobra.Command {
	minstCfg := startup.NewMinstConfig("", false, false) // empty cfg initialization

	// main cmd
	var rootCmd = &cobra.Command{
		Use:   "minst",
		Short: "Minst is a tool that lets you install all your software with a single command.",
		Long: "Minst is a configuration tool that reads your yml configuration file in order to \n" +
			"download and install all the software you want with a single command.",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help() // main command just runs help
		},
	}
	rootCmd.SetVersionTemplate("Minst version is 1.0.0\n")

	// cli commands wire up
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newInstallCmd(minstCfg))

	return rootCmd
}
