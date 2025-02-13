package cmd

import (
	"context"

	"github.com/bluewave-labs/checkmate-cli/internal/api/docker"
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a docker volume",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dockerClient, err := docker.NewDockerClient()

		if err != nil {
			logger.Error(err.Error())
		}

		defer dockerClient.Client.Close()

		// volumeName should be the first argument
		// backupPath should be the second argument
		volumeName := args[0]
		backupPath := args[1]

		err = dockerClient.BackupVolume(context.Background(), volumeName, backupPath)
		if err != nil {
			logger.Error(err.Error())
		}

		cmd.Println("Volume backup created successfully")
	},
}
