package models

// InstallFile - represents a fully-parsed yml file with user software
type InstallFile struct {
	Version          string `yaml:"version"`
	InstallProviders bool   `yaml:"install_providers"`

	Providers []struct {
		Name                    string   `yaml:"name"`
		BaseCommand             string   `yaml:"base_command"`
		InstallProviderCommands []string `yaml:"install_provider_commands"`
	} `yaml:"providers"`

	MachineSetup struct {
		ContinueOnError bool `yaml:"continue_on_error"`
		Installations   []struct {
			Provider string   `yaml:"provider"`
			Apps     []string `yaml:"apps"`
		} `yaml:"installations"`
	} `yaml:"machine_setup"`
}
