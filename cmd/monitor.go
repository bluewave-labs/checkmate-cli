package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bluewave-labs/checkmate-cli/internal/api/checkmate"
	checkmateTypes "github.com/bluewave-labs/checkmate-cli/internal/api/checkmate/types"
	"github.com/bluewave-labs/checkmate-cli/internal/cli/input"
	"github.com/bluewave-labs/checkmate-cli/internal/cli/output/visualizer"
	"github.com/bluewave-labs/checkmate-cli/internal/config"
	"github.com/bluewave-labs/checkmate-cli/internal/fs"
	"github.com/bluewave-labs/checkmate-cli/pkg/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var appConfig = config.AppConfig

var checkmateClient = checkmate.NewCheckmateClient(
	&appConfig,
	&http.Client{
		Timeout: 10 * time.Second,
	},
	checkmate.NewBearerAuthenticator(appConfig.APIKey),
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Manage your monitors",
}

var addMonitorCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var removeMonitorCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an existing monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listMonitorCmd = &cobra.Command{
	Use:   "list",
	Short: "List all monitors",
}

var listUpMonitorCmd = &cobra.Command{
	Use:   "up",
	Short: "List all up monitors",
	Run:   func(cmd *cobra.Command, args []string) {},
}
var listDownMonitorCmd = &cobra.Command{
	Use:   "down",
	Short: "List all down monitors",
	Run:   func(cmd *cobra.Command, args []string) {},
}
var listPausedMonitorCmd = &cobra.Command{
	Use:   "paused",
	Short: "List all paused monitors",
	Run:   func(cmd *cobra.Command, args []string) {},
}
var listAllMonitorCmd = &cobra.Command{
	Use:   "all",
	Short: "List all monitors",
	Run: func(cmd *cobra.Command, args []string) {
		// Command output format
		// "table" or "template"
		outputFormat, _ := cmd.Flags().GetString("output")

		// Get all monitors from the Checkmate API
		resp, err := checkmateClient.GetAllMonitors()
		if err != nil {
			logger.Error("error getting all monitors: " + err.Error())
		}

		// Parse the response data
		var allMonitors []checkmateTypes.Monitor
		for _, monitorData := range resp.Data.([]interface{}) {
			allMonitors = append(allMonitors, checkmateTypes.Monitor{
				Name:     monitorData.(map[string]interface{})["name"].(string),
				URL:      monitorData.(map[string]interface{})["url"].(string),
				Type:     checkmateTypes.MonitorType(monitorData.(map[string]interface{})["type"].(string)),
				UserID:   appConfig.UserID,
				TeamID:   appConfig.TeamID,
				IsActive: monitorData.(map[string]interface{})["isActive"].(bool),
				Status:   monitorData.(map[string]interface{})["status"].(bool),
			})
		}

		var upMonitors []checkmateTypes.Monitor
		var downMonitors []checkmateTypes.Monitor
		var pausedMonitors []checkmateTypes.Monitor

		for _, monitor := range allMonitors {
			if !monitor.IsActive {
				pausedMonitors = append(pausedMonitors, monitor)
			}

			switch monitor.Status {
			case true:
				upMonitors = append(upMonitors, monitor)
			case false:
				downMonitors = append(downMonitors, monitor)
			}
		}

		user := checkmateTypes.MonitorTemplate{
			InstanceURL:    appConfig.APIBaseURL,
			TeamID:         appConfig.TeamID,
			TotalMonitors:  getNameKeys(allMonitors), // Use the keys of the monitor list. It's a walkaround for using string slice in template.
			UpMonitors:     upMonitors,
			DownMonitors:   downMonitors,
			PausedMonitors: pausedMonitors,
		}

		var template visualizer.Visualizer

		if outputFormat == "table" {
			template = visualizer.Table{
				Header: []any{"Name", "URL", "Type", "Status"},
				Data:   user.MonitorTable(),
			}
		} else if outputFormat == "template" {
			template = visualizer.Template{
				Name: "monitor_list",
				Data: user,
			}
		} else {
			template = visualizer.Table{
				Header: []any{"Name", "URL", "Type", "Status"},
				Data:   user.MonitorTable(),
			}
		}

		err = template.Stdout()
		if err != nil {
			logger.Error("error rendering monitor list: " + err.Error())
		}
	},
}

var bulkImportMonitorCmd = &cobra.Command{
	Use:   "bulk-import",
	Short: "Bulk import monitors",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		monitorsFilePath := args[0]                  // Path to the file
		printLimit, _ := cmd.Flags().GetInt("limit") // Limit the number of monitors to be printed

		monitorCSV, err := fs.ReadCSVFile(monitorsFilePath)
		if err != nil {
			logger.Error(err.Error())
		}

		fmt.Printf("Monitors to be imported:\n\n")

		// Listing
		monitorNames, err := monitorCSV.ListColumn(0)
		if err != nil {
			logger.Error("error listing monitor names: " + err.Error())
		}

		for i := 0; i < monitorCSV.RowsCount; i++ {
			if i >= printLimit {
				fmt.Printf("...\nTotal: %s\n", color.BlueString(strconv.Itoa(monitorCSV.RowsCount)))
				break
			}
			color.Green(monitorNames[i] + "\n")
		}
		fmt.Println("---")

		// Confirmation
		userInput, err := input.StdIn("Do you want to continue? (y/n): ")
		if err != nil {
			logger.Error(err.Error())
		}

		if strings.ToLower(userInput) == "y" || strings.ToLower(userInput) == "yes" {
			fmt.Println("Importing monitors...")

			var monitors []checkmateTypes.Monitor

			for _, monitor := range monitorCSV.Rows {
				monitors = append(monitors, checkmateTypes.Monitor{
					UserID: appConfig.UserID,
					TeamID: appConfig.TeamID,
					Name:   monitor[0],
					URL:    monitor[1],
					Type:   checkmateTypes.MonitorType(monitor[2]),
				})
			}

			response, err := checkmateClient.CreateBulkMonitors(monitors)
			if err != nil {
				logger.Error("error creating bulk monitors: " + err.Error())
			}

			if response.Success {
				fmt.Println("Monitors imported successfully")
			} else {
				fmt.Println("Failed to import monitors")
			}

		} else {
			fmt.Println("Aborted")
		}
	},
}

func init() {
	bulkImportMonitorCmd.Flags().Int("limit", 5, "Limit the number of monitors to be printed")
	listAllMonitorCmd.Flags().StringP("output", "o", "table", "Output format (table, template)")
	listMonitorCmd.AddCommand(listUpMonitorCmd, listDownMonitorCmd, listPausedMonitorCmd, listAllMonitorCmd)
	monitorCmd.AddCommand(addMonitorCmd, removeMonitorCmd, listMonitorCmd, bulkImportMonitorCmd)
}

// FIXME: This function is in the wrong place.
func getNameKeys(monitorList []checkmateTypes.Monitor) []string {
	var names []string
	for _, monitor := range monitorList {
		names = append(names, monitor.Name)
	}
	return names
}
