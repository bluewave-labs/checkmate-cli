package output

import (
	"fmt"
	"strconv"
	"text/template"

	"github.com/fatih/color"
)

var TemplateFunctions = template.FuncMap{
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
}
