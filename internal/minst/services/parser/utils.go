package parser

import (
	"strings"

	"github.com/IgooorGP/minst/internal/minst/services/cmdrunner"
)

// ExtractCommandAndArgsFromString - reads a string and parses it to a terminal command
func ExtractCommandAndArgsFromString(command string, sep string) cmdrunner.TerminalCommand {
	var mainCommand string
	var commandArgs []string

	splitCommandString := strings.Split(command, sep)
	mainCommand = splitCommandString[0]

	if len(splitCommandString) > 1 {
		commandArgs = splitCommandString[1:]
	}

	return cmdrunner.TerminalCommand{Command: mainCommand, CommandArgs: commandArgs}
}

// ExtractCommandsAndArgsFromStrings - reads many strings and parses it to several terminal commands
func ExtractCommandsAndArgsFromStrings(commands []string, sep string) []cmdrunner.TerminalCommand {
	var terminalCommands []cmdrunner.TerminalCommand

	// provider's own install command
	for _, rawCommand := range commands {
		installCommand := ExtractCommandAndArgsFromString(rawCommand, " ")
		terminalCommands = append(terminalCommands, installCommand)
	}

	return terminalCommands
}
