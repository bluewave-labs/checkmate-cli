package fs

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSV struct {
	Columns      []string   // Column names
	Rows         [][]string // Rows excluding the first row
	RawData      [][]string // Raw data from the csv file
	ColumnsCount int        // Number of columns in the csv file
	RowsCount    int        // Excluding the first row which is the column names
	Shape        [2]int     // Shape of the CSV file, ColumnsCount:RowsCount
}

type CSVReader interface {
	ListColumn(index int) ([]string, error)

	// You can start the index from 0 for
	ListRow(index int) ([]string, error)
}

func (c *CSV) ListColumn(index int) ([]string, error) {
	if index < 0 || index >= len(c.Columns) {
		return nil, fmt.Errorf("index out of range")
	}

	var columnData []string
	for _, row := range c.Rows {
		columnData = append(columnData, row[index])
	}
	return columnData, nil
}

func (c *CSV) ListRow(index int) ([]string, error) {
	if index < 0 || index >= len(c.Rows) {
		return nil, fmt.Errorf("index out of range")
	}
	return c.Rows[index], nil
}

func ReadCSVFile(filePath string) (*CSV, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	// Read csv values using csv.Reader interface
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv data: %w", err)
	}

	csvFile := &CSV{
		RawData:      data,
		Columns:      data[0],
		Rows:         data[1:],
		ColumnsCount: len(data[0]),
		RowsCount:    len(data) - 1, // Excluding the first row which is the column names
		Shape:        [2]int{len(data[0]), len(data) - 1},
	}
	return csvFile, nil
}
