package cmdrunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldExtractStringFromCommandAndArgs(t *testing.T) {
	// arrange
	cmd := TerminalCommand{Command: "ls", CommandArgs: []string{"-ls", "."}}

	// act
	commandAndArgsString := ExtractStringFromCommandAndArgs(cmd.Command, cmd.CommandArgs)

	// assert
	expectedCommandAndArgsString := "ls -ls ."

	assert.Equal(t, expectedCommandAndArgsString, commandAndArgsString)
}

func TestShouldExtractStringFromCommandAndArgsForCommandWithoutArgs(t *testing.T) {
	// arrange
	cmd := TerminalCommand{Command: "ls", CommandArgs: nil}

	// act
	commandAndArgsString := ExtractStringFromCommandAndArgs(cmd.Command, cmd.CommandArgs)

	// assert
	expectedCommandAndArgsString := "ls"

	assert.Equal(t, expectedCommandAndArgsString, commandAndArgsString)
}
