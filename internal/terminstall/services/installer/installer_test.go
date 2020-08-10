package installer

import (
	"testing"

	"github.com/IgooorGP/terminstall/internal/terminstall/services/parser"
	"github.com/stretchr/testify/assert"
)

const testDataDirectory = "../../../../test/data/test_install_v1.yml"

func TestShouldGetProvidersForInstallFile(t *testing.T) {
	// arrange
	installFile := parser.ParseInstallFile(testDataDirectory)

	// act
	providers := ReadProvidersFromInstallFile(installFile)

	// assert
	assert.Len(t, providers, 2) // 2 providers
	assert.Equal(t, providers[0].Name, "brew")
	assert.Equal(t, providers[0].BaseCommand, "brew")

	assert.Equal(t, providers[1].Name, "brewcask")
	assert.Equal(t, providers[1].BaseCommand, "brew cask")
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
}
