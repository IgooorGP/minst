package main

import (
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Global configuration
var minstCfg *startup.MinstConfig

func main() {
	startup.SetupApp()
	minstCfg = startup.NewMinstConfig("", false, false) // empty cfg initialization

	// cli commands
	rootCmd := newMainCmd()
	installSubCmd := newInstallCmd(minstCfg)

	// cli commands wire up
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(installSubCmd)

	// exec the cli
	rootCmd.Execute()
}

func newMainCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "minst [COMMAND] [FLAGS]",
		Short: "Minst is a tool that lets you install all your software with a single command.",
		Long: "Minst is a configuration tool that reads your yml configuration file in order to \n" +
			"download and install all the software you want with a single command.",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msgf("Running main cmd")
		},
	}
	rootCmd.SetVersionTemplate("Minst version is 1.0.0\n")

	return rootCmd
}

// func main() {
// 	startup.SetupApp()

// 	log.Debug().Msgf("minster is parsing user-supplied commands...")
// 	commandLineArgs := parser.ParseCommandLineArgs()

// 	log.Info().Msgf("minster: using install file at %s ...", commandLineArgs.FilePath)
// 	installFile := parser.ParseInstallFile(commandLineArgs.FilePath)
// 	providers := parser.ReadProvidersFromInstallFile(installFile)

// 	for _, provider := range providers {
// 		log.Info().Msgf("Realizing setup for provider: %s", provider.Name)
// 		cmdrunner.RunCommands(provider.InstallProviderCommands, installFile.MachineSetup.ContinueOnError)

// 		log.Info().Msgf("Installing apps related to the provider...")
// 		cmdrunner.RunCommands(provider.InstallAppsCommands, installFile.MachineSetup.ContinueOnError)
// 	}

// 	log.Info().Msg("minster has finished its job!")
// }
