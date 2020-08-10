package cmdrunner

// ExtractStringFromCommandAndArgs - opposite way around: builds a single string from cmd its args
func ExtractStringFromCommandAndArgs(command string, args []string) string {
	commandAndArgsString := command

	for _, arg := range args {
		commandAndArgsString += (" " + arg)
	}

	return commandAndArgsString
}
