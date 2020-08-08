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
