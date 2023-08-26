package uploader

import (
	"encoding/csv"
	"os"
)

type csvWriter struct {
	writer *csv.Writer
}

func newCsvWriter(comma rune) *csvWriter {
	writer := csv.NewWriter(os.Stdout)

	writer.Comma = comma

	err := writer.Write(outData{}.titles())
	if err != nil {
		printErr("csv", err)
	}

	return &csvWriter{
		writer: writer,
	}
}

func (c *csvWriter) write(od outData) {
	defer c.writer.Flush()

	err := c.writer.Write(od.toStrings())
	if err != nil {
		printErr(od.Source, err)
	}
}
