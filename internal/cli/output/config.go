package output

import (
	"github.com/fatih/color"
)

var ConfigOverwriteWarning = "This action will" +
	" " +
	color.RedString("overwrite") +
	" " +
	"the current configuration."

var ConfigSetAPIKeyMessage = "Enter your Checkmate API Key:"
var ConfigSetUserIDMessage = "Enter your Checkmate User ID:"
var ConfigSetTeamIDMessage = "Enter your Checkmate Team ID:"
var ConfigSetBaseURLMessage = "Enter your Checkmate API Base URL:"
