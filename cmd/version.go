package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of Checkmate CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Checkmate CLI Version: %s\n", Version)
	},
}
