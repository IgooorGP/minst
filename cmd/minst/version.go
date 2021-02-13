package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// VersionCmd - minst version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of Hugo",
	Long:  `All software have versions. This is Minst's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msgf("Minst Version 1.0.0")
	},
}
