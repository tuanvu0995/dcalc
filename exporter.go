package main

import (
	"encoding/csv"
	"os"
)

type ExportType string

const (
	ExportTypeCSV ExportType = "csv"
)

type Exporter struct {
	App        *App
	ExportType ExportType
	Records    [][]string
}

func NewExporter(app *App) *Exporter {
	return &Exporter{
		App: app,
	}
}

func (e *Exporter) Data() *Exporter {
	e.Records = e.App.GetRecords()
	return e
}

func (e *Exporter) ToCSV(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"name", "op", "calc", "value"}
	writer.Write(headers)
	for _, record := range e.Records {
		writer.Write(record)
	}
}
