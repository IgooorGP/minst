/*
Main entry-point of the Minst "My-installer" application.
*/

package main

import (
	"github.com/IgooorGP/minst/internal/minst/startup"
	"github.com/IgooorGP/minst/pkg/cmd"
)

func main() {
	startup.SetupApp()

	// cli commands
	minstCmd := cmd.NewMinstCmd()

	// exec the cli
	minstCmd.Execute()
}
