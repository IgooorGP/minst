package cmd

import (
	"github.com/IgooorGP/minst/internal/minst/services/cmdrunner"
	"github.com/IgooorGP/minst/internal/minst/services/parser"
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newInstallCmd creates a new install sub command and add its flags
func newInstallCmd(cfg *startup.MinstConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs all software found in an Minstall file.",
		Run:   installCmdHandlerFactory(cfg),
	}

	cmd.PersistentFlags().StringVarP(&cfg.InstallFilePath, "file", "f", "minstall.yml", "The install file path")
	cmd.PersistentFlags().BoolVar(&cfg.DryRun, "dry-run", false, "Runs the program without installing or downloading anything")
	cmd.PersistentFlags().BoolVar(&cfg.NoParallelInstalls, "no-parallel-installs", false, "Installs each software in a sequence and not in parallel")

	return cmd
}

// factory for creating install cmds - adds minstCfg to the closure lexical scope
func installCmdHandlerFactory(cfg *startup.MinstConfig) func(cmd *cobra.Command, args []string) {

	return func(cmd *cobra.Command, args []string) {
		log.Info().Msgf("minster: using install file at %s ...", cfg.InstallFilePath)

		installFile := parser.ParseInstallFile(cfg.InstallFilePath)
		softwareProviders := parser.CreateSoftwareProviders(installFile)

		for _, provider := range softwareProviders {

			for _, cmdStr := range provider.InstallProviderCommands {
				installCmd := cmdrunner.ExtractStringFromCommandAndArgs(cmdStr.Command, cmdStr.CommandArgs)
				log.Info().Msg(installCmd)
			}

			for _, cmdStr := range provider.InstallAppsCommands {
				installCmd := cmdrunner.ExtractStringFromCommandAndArgs(cmdStr.Command, cmdStr.CommandArgs)
				log.Info().Msg(installCmd)
			}
			// cmdrunner.RunCommands(provider.InstallProviderCommands, installFile.MachineSetup.ContinueOnError)

			// // provider.InstallAppsCommands -> ["brew install slack", "brew install chrome", ...]
			// // here we either run in sequence or in parallel with goroutines

			// cmdrunner.RunCommands(provider.InstallAppsCommands, installFile.MachineSetup.ContinueOnError)
		}
	}
}
