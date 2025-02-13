package cmd

import (
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/spf13/cobra"
)

var Version string = "development"

var rootCmd = &cobra.Command{
	Use:   "checkmate",
	Short: "Checkmate CLI is a tool for managing monitors",
	Long: `Checkmate CLI is a command line interface for managing monitors.
It allows you to add, remove, and list monitors with ease.`,
	Version: Version,
}

func init() {
	rootCmd.SetVersionTemplate("Checkmate CLI Version: {{.Version}}\n")
	rootCmd.AddCommand(versionCmd, monitorCmd, configCmd, backupCmd, restoreCmd)
}

// Execute executes the root command of the CLI application.
// It invokes the command execution process and handles any encountered errors
// by logging a fatal error and terminating the application.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		logger.Error(err.Error())
	}
}
