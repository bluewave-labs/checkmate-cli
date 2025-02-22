package visualizer

import (
	"cmp"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	defaultHeaderColor = color.New(color.FgWhite, color.Underline)
	defaultColumnColor = color.New(color.FgYellow)
)

type Table struct {
	Header      []interface{} // Header of the table
	Data        [][]string    // Data of the table
	HeaderColor *color.Color  // Color of the header
	ColumnColor *color.Color  // Color of the columns
}

func (t Table) Stdout() error {
	tbl := table.New(t.Header...)
	t.tableFormat(tbl, nil, nil)
	tbl.WithWriter(os.Stdout)
	tbl.SetRows(t.Data)
	tbl.Print()
	return nil
}

func (t *Table) tableFormat(table table.Table, headerColor *color.Color, columnColor *color.Color) {
	headerColor = cmp.Or(headerColor, defaultHeaderColor)
	columnColor = cmp.Or(columnColor, defaultColumnColor)

	headerFormat := headerColor.SprintfFunc()
	columnFormat := columnColor.SprintfFunc()
	table.WithHeaderFormatter(headerFormat).WithFirstColumnFormatter(columnFormat)
}
