package models

// TerminalCommand - Abstraction of a terminal command that can be executed with or without errors
type TerminalCommand struct {
	Command      string
	CommandArgs  []string
	Executed     bool
	HasError     bool
	ErrorMessage string
}
