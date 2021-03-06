package cmdrunner

import (
	"bytes"
	"os/exec"

	"github.com/rs/zerolog/log"
)

// TerminalCommand - Abstraction of a terminal command that can be executed with or without errors
type TerminalCommand struct {
	Command      string
	CommandArgs  []string
	Executed     bool
	HasError     bool
	ErrorMessage string
}

// RunCommand - invokes a shell command and prints its stdout (or stderr if an error happens)
func RunCommand(terminalCommand TerminalCommand, continueOnError bool) TerminalCommand {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// builds the command and gets references to its stdout and stderr
	log.Info().Msgf("[COMMAND]: %s", ExtractStringFromCommandAndArgs(terminalCommand.Command, terminalCommand.CommandArgs))
	commmand := exec.Command(terminalCommand.Command, terminalCommand.CommandArgs...)
	commmand.Stdout = &stdout
	commmand.Stderr = &stderr

	// invokes the command till completion
	err := commmand.Run()
	terminalCommand.Executed = true
	executedTerminalCommand := TerminalCommand{Command: terminalCommand.Command, CommandArgs: terminalCommand.CommandArgs, Executed: true}

	if err != nil && !continueOnError {
		log.Fatal().Msgf("[COMMAND_RESULT]: execution failed and won't continue on errors! process stderr: %s", stderr.String())
	}

	if err != nil && continueOnError {
		executedTerminalCommand.HasError = true
		executedTerminalCommand.ErrorMessage = stderr.String()

		log.Error().Msgf("[COMMAND_RESULT]: execution failed! stderr: %s", executedTerminalCommand.ErrorMessage)
		return executedTerminalCommand
	}

	// even if no sys errors were raised, e.g. return 1, stderr can still have been used for warnings!
	if stderr.String() != "" {
		log.Info().Msgf("[COMMAND_RESULT]: execution raised warnings! stderr: %s", stderr.String())
	} else {
		log.Info().Msgf("[COMMAND_RESULT]: executed successfully!\n")
	}

	return executedTerminalCommand
}

// RunCommands - Synchronously invokes a list of terminal commands.
func RunCommands(terminalCommands []TerminalCommand, continueOnError bool) []TerminalCommand {
	var executedTerminalCommands []TerminalCommand

	for _, terminalCommand := range terminalCommands {
		executedTerminalCommands = append(executedTerminalCommands, RunCommand(terminalCommand, continueOnError))
	}

	return executedTerminalCommands
}

// RunCommandsAsync - Asynchronously invokes a list of terminal commands by using a worker pool of a given size.
func RunCommandsAsync(commands []TerminalCommand, continueOnError bool, numberOfWorkers int) {

}
