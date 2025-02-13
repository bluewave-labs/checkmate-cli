package output

import (
	"github.com/fatih/color"
)

var ConfigSetOverwriteConfirmMessage = "This action will" +
	" " +
	color.RedString("overwrite") +
	" " +
	"the current \"api_key\" in the config file\nEnter your Checkmate API Key:"
