package main

import (
	"github.com/IgooorGP/terminstall/internal/terminstall/startup"
	"github.com/rs/zerolog/log"
)

func main() {
	startup.SetupApp()
	log.Info().Msg("hello world!")
}
