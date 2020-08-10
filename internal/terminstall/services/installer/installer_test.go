package installer

import (
	"testing"

	"github.com/IgooorGP/terminstall/internal/terminstall/services/parser"
	"github.com/stretchr/testify/assert"
)

const testDataDirectory = "../../../../test/data/test_install_less_apps.yaml"

func TestShouldGetProvidersForInstallFile(t *testing.T) {
	// arrange
	installFile := parser.ParseInstallFile(testDataDirectory)

	// act
	providers := ReadProvidersFromInstallFile(installFile)

	// assert
	assert.Len(t, providers, 2) // 2 providers
	assert.Equal(t, providers[0].Name, "brew")
	assert.Equal(t, providers[0].BaseCommand, "brew install")

	assert.Equal(t, providers[1].Name, "brewcask")
	assert.Equal(t, providers[1].BaseCommand, "brew cask install")
}

func TestGetInstallCommandsForProvider(t *testing.T) {
	// arrange
	installFile := parser.ParseInstallFile(testDataDirectory)

	// act
	providers := ReadProvidersFromInstallFile(installFile)
	installCommands := GetInstallCommandsForProvider(installFile, providers[1]) // brew cask
	brewCaskProvider := providers[1]

	// assert
	assert.Equal(t, brewCaskProvider.Name, "brewcask")
	assert.Len(t, installCommands, 2)

	// assert - for brew cask install slack
	assert.Equal(t, installCommands[0].Command, "brew")
	assert.Equal(t, installCommands[0].CommandArgs, []string{"cask", "install", "slack"})
	assert.Equal(t, installCommands[0].Executed, false)
	assert.Equal(t, installCommands[0].HasError, false)

	// assert - for brew cask install vscode
	assert.Equal(t, installCommands[1].Command, "brew")
	assert.Equal(t, installCommands[1].CommandArgs, []string{"cask", "install", "vscode"})
	assert.Equal(t, installCommands[1].Executed, false)
	assert.Equal(t, installCommands[1].HasError, false)
}
