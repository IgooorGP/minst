package parser

import (
	"strings"

	"github.com/IgooorGP/minst/internal/minst/models"
)

// ExtractCommandAndArgsFromString - reads a string and parses it to a terminal command
func ExtractCommandAndArgsFromString(command string, sep string) models.TerminalCommand {
	var mainCommand string
	var commandArgs []string

	splitCommandString := strings.Split(command, sep)
	mainCommand = splitCommandString[0]

	if len(splitCommandString) > 1 {
		commandArgs = splitCommandString[1:]
	}

	return models.TerminalCommand{Command: mainCommand, CommandArgs: commandArgs}
}

// ExtractCommandsAndArgsFromStrings - reads many strings and parses it to several terminal commands
func ExtractCommandsAndArgsFromStrings(commands []string, sep string) []models.TerminalCommand {
	var terminalCommands []models.TerminalCommand

	// provider's own install command
	for _, rawCommand := range commands {
		installCommand := ExtractCommandAndArgsFromString(rawCommand, " ")
		terminalCommands = append(terminalCommands, installCommand)
	}

	return terminalCommands
}
