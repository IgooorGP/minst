package parser

import (
	"io/ioutil"
	"path/filepath"

	"github.com/IgooorGP/minst/internal/minst/models"
	"github.com/goccy/go-yaml"
)

// ParseInstallFile - parses the install file
func ParseInstallFile(file string) models.InstallFile {
	var installFile models.InstallFile

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
