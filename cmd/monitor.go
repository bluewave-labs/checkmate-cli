package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/bluewave-labs/checkmate-cli/internal/cli"
	"github.com/bluewave-labs/checkmate-cli/internal/config"
	"github.com/bluewave-labs/checkmate-cli/internal/fs"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		teamId := "team-123"
		instanceURL := config.AppConfig.APIBaseURL

		monitorListTemplate := `Your Checkmate Monitors on instance: %s
Team ID: %s

Total Monitors : %s
		`
		monitorListOutput := fmt.Sprintf(monitorListTemplate,
			color.GreenString(instanceURL), // Your instance URL
			color.CyanString(teamId),       // Your team ID

			color.BlueString("10"), // Total Monitors
		)

		if len(args) > 0 {
			switch args[0] {
			case "up":
				upListTemplate := `
Up Monitors    : %s
	%s
`
				upMonitors := `- Google
	- Facebook
	- Twitter
	- GitHub
	- Bluesky Social`

				monitorListOutput += fmt.Sprintf(upListTemplate,
					color.GreenString("5"), // Up Monitors
					upMonitors,             // Indented list of up monitors
				)
			case "down":
				downListTemplate := `
Down Monitors  : %s
	%s`
				downMonitors := `- Google 404
	- LinkedIn
	- Netflix`

				monitorListOutput += fmt.Sprintf(downListTemplate,
					color.RedString("3"), // Down Monitors
					downMonitors,         // Indented list of down monitors
				)
			case "paused":
				pausedListTemplate := `
Paused Monitors: %s
	%s`
				pausedMonitors := `- LinkedIn
	- Instagram`
				monitorListOutput += fmt.Sprintf(pausedListTemplate,
					color.YellowString("2"), // Paused Monitors
					pausedMonitors,          // Indented list of paused monitors
				)
			case "another":
				var tmplFile = "template/monitor_list.tmpl"
				tmpl, err := template.New("monitor_list.tmpl").Funcs(template.FuncMap{
					"green": func(input string) string {
						return color.GreenString(input)
					},
					"red": func(input string) string {
						return color.RedString(input)
					},
					"yellow": func(input string) string {
						return color.YellowString(input)
					},
					"blue": func(input string) string {
						return color.BlueString(input)
					},
					"cyan": func(input string) string {
						return color.CyanString(input)
					},
					"join": func(input []string) string {
						return fmt.Sprintf("%v", input)
					},
					"intToString": func(i int) string {
						return strconv.Itoa(i)
					},
				}).ParseFiles(tmplFile)

				if err != nil {
					panic("Error parsing template: " + err.Error())
				}
				type MonitorTemplate struct {
					InstanceURL    string
					TeamID         string
					TotalMonitors  []string
					UpMonitors     []string
					DownMonitors   []string
					PausedMonitors []string
				}
				user := MonitorTemplate{
					InstanceURL:    instanceURL,
					TeamID:         teamId,
					TotalMonitors:  []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
					UpMonitors:     []string{"a", "b", "c", "d", "e"},
					DownMonitors:   []string{"a", "b", "c"},
					PausedMonitors: []string{"a", "b"},
				}

				err = tmpl.Execute(os.Stdout, user)
				if err != nil {
					panic(err)
				}
			default:
				// monitorListOutput += upMonitor
				// monitorListOutput += downMonitor
				// monitorListOutput += pausedMonitor
			}
		}

		// api request then
		//
		// response.data.summary
		// totalMonitors int
		// upMonitors int
		// downMonitors int
		// pausedMonitors int

		fmt.Println(monitorListOutput)
	},
}

var bulkImportMonitorCmd = &cobra.Command{
	Use:   "bulk-import",
	Short: "Bulk import monitors",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		monitors := args[0] // Path to the file
		printLimit, _ := cmd.Flags().GetInt("limit")

		monitorCSV, err := fs.ReadCSVFile(monitors)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Printf("Monitors to be imported:\n\n")

		monitorNames, err := monitorCSV.ListColumn(0)
		if err != nil {
			log.Fatalln(err)
		}

		for i := 0; i < monitorCSV.RowsCount; i++ {
			if i >= printLimit {
				fmt.Printf("...\nTotal: %s\n", color.BlueString(strconv.Itoa(monitorCSV.RowsCount)))
				break
			}
			color.Green(monitorNames[i] + "\n")
		}
		fmt.Println("---")

		userInput, err := cli.StdIn("Do you want to continue? (y/n): ")

		if err != nil {
			log.Fatalln(err)
		}

		if strings.ToLower(userInput) == "y" {
			fmt.Println("Importing monitors...")
			// CHECKMATE API CALL
		} else {
			fmt.Println("Aborted")
		}
	},
}

func init() {
	bulkImportMonitorCmd.Flags().Int("limit", 5, "Limit the number of monitors to be printed")
	monitorCmd.AddCommand(addMonitorCmd, removeMonitorCmd, listMonitorCmd, bulkImportMonitorCmd)
}
