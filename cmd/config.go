package cmd

import (
	"fmt"
	"log"

	"github.com/bluewave-labs/checkmate-cli/internal/cli"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the current configuration",
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := fmt.Sprintf(
			"This action will %s the current \"api_key\" in the config file\nEnter your Checkmate API Key: ",
			color.RedString("overwrite"),
		)

		// Get auth_token from os.Stdin
		apiKey, err := cli.StdIn(prompt)
		if err != nil {
			log.Fatalf("error reading input: %v\n", err)
		}

		// Set the auth_token in viper
		viper.Set("api_key", apiKey)

		// Save config to file
		err = viper.WriteConfig()
		if err != nil {
			log.Fatalf("error saving config: %v\n", err)
		}

		cmd.Println("Checkmate token saved successfully!")
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
