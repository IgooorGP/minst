package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDataDirectory = "../../../../test/data/test_install_v1.yml"

func TestInstallFileParser(t *testing.T) {
	// act
	parsedInstallFile := ParseInstallFile(testDataDirectory)

	// assert -- providers
	assert.True(t, parsedInstallFile.InstallProviders)
	assert.Equal(t, len(parsedInstallFile.Providers), 2)

	assert.Equal(t, parsedInstallFile.Providers[0].Name, "brew")
	assert.Equal(t, parsedInstallFile.Providers[0].InstallAppsCommand, "brew")
	assert.Equal(t, parsedInstallFile.Providers[0].InstallProviderCommands, []string{
		"/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)\"", "brew update"})

	assert.Equal(t, parsedInstallFile.Providers[1].Name, "brewcask")
	assert.Equal(t, parsedInstallFile.Providers[1].InstallAppsCommand, "brew cask")

	assert.True(t, parsedInstallFile.MachineSetup.ContinueOnError)
	assert.Len(t, parsedInstallFile.MachineSetup.Installations, 2)

	assert.Equal(t, parsedInstallFile.MachineSetup.Installations[0].Provider, "brewcask")
	assert.Equal(t, parsedInstallFile.MachineSetup.Installations[0].Apps, []string{"slack", "alfred", "vscode", "pycharm"})

	assert.Equal(t, parsedInstallFile.MachineSetup.Installations[1].Provider, "brew")
	assert.Equal(t, parsedInstallFile.MachineSetup.Installations[1].Apps, []string{"aws", "gcloud", "docker", "kubectl", "k9s"})
}
