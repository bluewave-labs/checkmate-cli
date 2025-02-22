package visualizer

import (
	"fmt"
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

	if err := tmpl.Execute(os.Stdout, t.Data); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}

	return nil
}
