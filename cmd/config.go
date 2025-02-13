package cmd

import (
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
		// Get api_key from os.Stdin
		apiKey, err := input.StdIn(output.ConfigSetOverwriteConfirmMessage)
		if err != nil {
			logger.Errorf("error reading input: %v\n", err)
		}

		// Set the api_key in viper
		viper.Set("api_key", apiKey)

		// Save config to file
		err = viper.WriteConfig()
		if err != nil {
			logger.Errorf("error saving config: %v\n", err)
		}

		cmd.Println("Checkmate token saved successfully!")
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
