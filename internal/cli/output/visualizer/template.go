package visualizer

import (
	"html/template"
	"os"

	"github.com/bluewave-labs/checkmate-cli/internal/cli/output"
)

type Template struct {
	Name string
	Data any
}

func (t Template) Stdout() error {
	tmpl, err := template.New(t.Name).Funcs(output.TemplateFunctions).Parse(output.MonitorListTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, t.Data)
}
