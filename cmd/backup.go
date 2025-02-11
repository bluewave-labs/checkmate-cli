package cmd

import (
	"context"
	"log"

	"github.com/bluewave-labs/checkmate-cli/internal/api/docker"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a docker volume",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dockerClient, err := docker.NewDockerClient()

		if err != nil {
			log.Fatalln("Error creating docker client", err)
		}

		defer dockerClient.Client.Close()

		// User input
		// volumeName should be the first argument
		// backupPath should be the second argument

		// PANIIKKKK
		volumeName := args[0]
		backupPath := args[1]

		err = dockerClient.BackupVolume(context.Background(), volumeName, backupPath)
		if err != nil {
			log.Fatalln("Error backing up volume", err)
		}

		cmd.Println("Volume backup created successfully")
	},
}
