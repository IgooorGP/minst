package parser

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

// CommandLineArgs - Contains all command line args supplied by the user
type CommandLineArgs struct {
	FilePath string
}

// InstallFile - represents a fully-parsed yml file
type InstallFile struct {
	Version          string `yaml:"version"`
	InstallProviders bool   `yaml:"install_providers"`

	Providers []struct {
		Name                    string   `yaml:"name"`
		InstallAppsCommand      string   `yaml:"install_apps_command"`
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

// ParseCommandLineArgs - parses command line args into a CommandLineArgs struct
func ParseCommandLineArgs() CommandLineArgs {
	installFilePath := flag.String("install-file", "./install.yml", "Filepath of the YML to be used to setup the installation.")
	flag.Parse()

	return CommandLineArgs{FilePath: *installFilePath}
}

// ParseInstallFile - parses the install file
func ParseInstallFile(file string) InstallFile {
	var installFile InstallFile
	filePath, _ := filepath.Abs(file)
	yamlFileBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFileBytes, &installFile)

	if err != nil {
		panic(err)
	}

	return installFile
}
