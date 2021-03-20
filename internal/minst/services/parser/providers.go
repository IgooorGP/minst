package parser

import (
	"github.com/IgooorGP/minst/internal/minst/models"
)

// CreateSoftwareProviders - Reads providers from install file along with its target commands
func CreateSoftwareProviders(installFile models.InstallFile) []models.SoftwareProvider {
	var providers []models.SoftwareProvider

	for _, providerData := range installFile.Providers {

		newProvider := models.SoftwareProvider{
			Name:                    providerData.Name,
			BaseCommand:             providerData.BaseCommand,
			InstallProviderCommands: ExtractCommandsAndArgsFromStrings(providerData.InstallProviderCommands, " "),
			InstallAppsCommands:     GetInstallCommandsForProvider(installFile, providerData.Name, providerData.BaseCommand),
		}

		providers = append(providers, newProvider)
	}

	return providers
}

// GetInstallCommandsForProvider - Gets desired apps and build install commands with the provider
func GetInstallCommandsForProvider(installFile models.InstallFile, providerName string, providerBaseCommand string) []models.TerminalCommand {
	var installAppCommandsString []string

	for _, providerInstallations := range installFile.MachineSetup.Installations {

		// finds the provider in the installation file
		if providerInstallations.Provider == providerName {
			// builds all install commands with the provider's base install command for each app
			for _, appName := range providerInstallations.Apps {
				installAppCommandsString = append(installAppCommandsString, providerBaseCommand+" "+appName)
			}
		}

	}

	// finally, convert the strings, "brew install slack" -> {Command: "brew install", Args: "slack", ...}
	return ExtractCommandsAndArgsFromStrings(installAppCommandsString, " ")
}
