package main

import (
	"github.com/IgooorGP/minst/internal/minst/services/cmdrunner"
	"github.com/IgooorGP/minst/internal/minst/services/parser"
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/rs/zerolog/log"
)

func main() {
	startup.SetupApp()

	log.Debug().Msgf("minster is parsing user-supplied commands...")
	commandLineArgs := parser.ParseCommandLineArgs()

	log.Info().Msgf("minster: using install file at %s ...", commandLineArgs.FilePath)
	installFile := parser.ParseInstallFile(commandLineArgs.FilePath)
	providers := parser.ReadProvidersFromInstallFile(installFile)

	for _, provider := range providers {
		log.Info().Msgf("Realizing setup for provider: %s", provider.Name)
		cmdrunner.RunCommands(provider.InstallProviderCommands, installFile.MachineSetup.ContinueOnError)

		log.Info().Msgf("Installing apps related to the provider...")
		cmdrunner.RunCommands(provider.InstallAppsCommands, installFile.MachineSetup.ContinueOnError)
	}

	log.Info().Msg("minster has finished its job!")
}
