package cmd

import (
	"context"

	"github.com/bluewave-labs/checkmate-cli/internal/api/docker"
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a volume from a backup",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dockerClient, err := docker.NewDockerClient()

		if err != nil {
			cmd.Println("Error creating docker client", err)
		}

		defer dockerClient.Client.Close()

		// User input
		// volumeName should be the first argument
		// backupPath should be the second argument

		// PANIIKKKK
		volumeName := args[0]
		backupPath := args[1]

		err = dockerClient.RestoreVolume(context.Background(), volumeName, backupPath)
		if err != nil {
			logger.Error(err.Error())
		}

		cmd.Println("Volume restored successfully")
	},
}
