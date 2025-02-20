package cmd

import (
	"fmt"

	"github.com/bluewave-labs/checkmate-cli/internal/cli/input"
	"github.com/bluewave-labs/checkmate-cli/internal/cli/output"
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use: "config",
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run: func(cmd *cobra.Command, args []string) {
		// Print the config overwrite warning
		fmt.Print(output.ConfigOverwriteWarning)

		// Get base_url from os.Stdin
		baseURL, err := input.StdIn(output.ConfigSetBaseURLMessage)
		if err != nil {
			logger.Errorf("error reading input: %v\n", err)
		}
		viper.Set("base_url", baseURL)

		// Get api_key from os.Stdin
		apiKey, err := input.StdIn(output.ConfigSetAPIKeyMessage)
		if err != nil {
			logger.Errorf("error reading input: %v\n", err)
		}
		viper.Set("api_key", apiKey)

		// Get user ID from os.Stdin
		userID, err := input.StdIn(output.ConfigSetUserIDMessage)
		if err != nil {
			logger.Errorf("error reading input: %v\n", err)
		}
		viper.Set("user_id", userID)

		// Get team ID from os.Stdin
		teamID, err := input.StdIn(output.ConfigSetTeamIDMessage)
		if err != nil {
			logger.Errorf("error reading input: %v\n", err)
		}
		viper.Set("team_id", teamID)

		// Save config to file
		err = viper.WriteConfig()
		if err != nil {
			logger.Errorf("error saving config: %v\n", err)
		}

		cmd.Println("Checkmate configuration updated")
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
