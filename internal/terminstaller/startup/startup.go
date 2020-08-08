package startup

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupApp - configures the application on app's start
func SetupApp() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zerolog.SetGlobalLevel(loggingLevelMap[loggingLevel])

	log.Logger = log.Output(output)
}
