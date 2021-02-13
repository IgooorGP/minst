package startup

import "github.com/rs/zerolog"

var loggingLevelMap = map[string]zerolog.Level{
	"panic": zerolog.PanicLevel,
	"fatal": zerolog.FatalLevel,
	"warn":  zerolog.WarnLevel,
	"info":  zerolog.InfoLevel,
	"debug": zerolog.DebugLevel,
}

// LoggingLevel - logging level of the application
const loggingLevel string = "info"

// MinstConfig tracks new Minst runtime config options
type MinstConfig struct {
	InstallFilePath    string
	DryRun             bool
	NoParallelInstalls bool
}

// NewMinstConfig compiles all cli config into a single struct: MinstConfig
func NewMinstConfig(installFilePath string, dryRun bool, noParallelInstalls bool) *MinstConfig {
	return &MinstConfig{
		InstallFilePath:    installFilePath,
		DryRun:             dryRun,
		NoParallelInstalls: noParallelInstalls,
	}
}

// InstallFile - represents a fully-parsed yml file
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
