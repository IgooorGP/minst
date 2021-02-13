package main

import (
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newInstallCmd creates a new install sub command to handle installs
func newInstallCmd(cfg *startup.MinstConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs all software found on the install file.",
		Run:   installCmdHandler,
	}
	cmd.PersistentFlags().StringVarP(&cfg.InstallFilePath, "file", "f", "minstall.yml", "The install file path")
	cmd.PersistentFlags().BoolVar(&cfg.DryRun, "dry-run", false, "Runs the program without installing or downloading anything")
	cmd.PersistentFlags().BoolVar(&cfg.NoParallelInstalls, "no-parallel-installs", false, "Installs each software in a sequence and not in parallel")

	return cmd
}

func installCmdHandler(cmd *cobra.Command, args []string) {
	log.Info().Msgf("installFilePath: %s", minstCfg.InstallFilePath)
	log.Info().Msgf("dryRun: %b", minstCfg.DryRun)
	log.Info().Msgf("noParallelInstalls: %b", minstCfg.NoParallelInstalls)
}
