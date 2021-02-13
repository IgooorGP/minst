package startup

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupApp - configures the application on app's start
func SetupApp() {
	output := zerolog.ConsoleWriter{Out: os.Stdout}

	// removes level and timestamp from std logging
	output.FormatLevel = func(i interface{}) string {
		return ""
	}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}

	zerolog.SetGlobalLevel(loggingLevelMap[loggingLevel])
	log.Logger = log.Output(output)
}
