package parser

import (
	"testing"

	"github.com/IgooorGP/minst/internal/minst/models"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateTerminalCommandWithArgs(t *testing.T) {
	// arrange
	command := "ls -ls ." // a command with args

	// act
	extractedCommand := ExtractCommandAndArgsFromString(command, " ")

	// assert
	expectedCommand := models.TerminalCommand{Command: "ls", CommandArgs: []string{"-ls", "."}}

	assert.Equal(t, expectedCommand, extractedCommand)
}

func TestShouldCreateTerminalCommandWithoutArgs(t *testing.T) {
	// arrange
	command := "ls" // a command without args

	// act
	extractedCommand := ExtractCommandAndArgsFromString(command, " ")

	// assert
	expectedCommand := models.TerminalCommand{Command: "ls", CommandArgs: nil}

	assert.Equal(t, expectedCommand, extractedCommand)
}

func TestShouldCreateTerminalCommandsFromCommandsList(t *testing.T) {
	// arrange
	commands := []string{"ls -ls .", "ls -ls /somefolder", "ps"} // a series of commands (some with args and some without)

	// act
	extractedCommands := ExtractCommandsAndArgsFromStrings(commands, " ")

	// assert
	expectedCommands := []models.TerminalCommand{
		{Command: "ls", CommandArgs: []string{"-ls", "."}},
		{Command: "ls", CommandArgs: []string{"-ls", "/somefolder"}},
		{Command: "ps", CommandArgs: nil},
	}

	assert.Equal(t, expectedCommands, extractedCommands)
}
