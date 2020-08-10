package installer

import (
	"github.com/IgooorGP/terminstall/internal/terminstall/services/parser"
)

// Provider - represents a software provider in a way that it's used to install such software
type Provider struct {
	Name                    string            // provider name
	BaseCommand             string            // base command to install each app
	InstallProviderCommands []TerminalCommand // commands to install this provider
	InstallAppsCommands     []TerminalCommand // base command + each app name
}

// ReadProvidersFromInstallFile - Reads providers from install file along with its target commands
func ReadProvidersFromInstallFile(installFile parser.InstallFile) []Provider {
	var providers []Provider

	for _, providerData := range installFile.Providers {
		providerInstallCommands := ExtractCommandsAndArgsFromStrings(providerData.InstallProviderCommands, " ")
		newProvider := Provider{Name: providerData.Name, InstallProviderCommands: providerInstallCommands, BaseCommand: providerData.BaseCommand}
		providers = append(providers, newProvider)
	}

	return providers
}

// GetInstallCommandsForProvider - Gets desired apps and build install commands with the provider
func GetInstallCommandsForProvider(installFile parser.InstallFile, provider Provider) []TerminalCommand {
	var installAppCommandsString []string

	for _, providerInstallations := range installFile.MachineSetup.Installations {
		// finds the provider in the installation file
		if providerInstallations.Provider == provider.Name {
			// builds all install commands with the provider's base install command for each app
			for _, appName := range providerInstallations.Apps {
				installAppCommandsString = append(installAppCommandsString, provider.BaseCommand+" "+appName)
			}
		}
	}

	return ExtractCommandsAndArgsFromStrings(installAppCommandsString, " ")
}
