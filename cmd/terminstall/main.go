package main

import (
	"github.com/IgooorGP/terminstall/internal/terminstall/services/cmdrunner"
	"github.com/IgooorGP/terminstall/internal/terminstall/services/parser"
	"github.com/IgooorGP/terminstall/internal/terminstall/startup"
	"github.com/rs/zerolog/log"
)

func main() {
	startup.SetupApp()

	log.Debug().Msgf("Terminstaller is parsing user-supplied commands...")
	commandLineArgs := parser.ParseCommandLineArgs()

	log.Info().Msgf("Terminstaller: using install file at %s ...", commandLineArgs.FilePath)
	installFile := parser.ParseInstallFile(commandLineArgs.FilePath)
	providers := parser.ReadProvidersFromInstallFile(installFile)

	for _, provider := range providers {
		log.Info().Msgf("Realizing setup for provider: %s", provider.Name)
		cmdrunner.RunCommands(provider.InstallProviderCommands, installFile.MachineSetup.ContinueOnError)

		log.Info().Msgf("Installing apps related to the provider...")
		cmdrunner.RunCommands(provider.InstallAppsCommands, installFile.MachineSetup.ContinueOnError)
	}

	log.Info().Msg("Terminstaller has finished its job!")
}
