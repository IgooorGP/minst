package models

// SoftwareProvider - represents a software provider in a way that it's used to install such software
type SoftwareProvider struct {
	Name                    string            // provider name
	BaseCommand             string            // base command to install each app
	InstallProviderCommands []TerminalCommand // commands to install this provider
	InstallAppsCommands     []TerminalCommand // base command + each app name
}
