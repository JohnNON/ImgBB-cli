package uploader

import (
	"encoding/json"
	"os"
)

type jsonWriter struct {
	writer *json.Encoder
}

func newJsonWriter() *jsonWriter {
	return &jsonWriter{
		writer: json.NewEncoder(os.Stdout),
	}
}

func (j *jsonWriter) write(od outData) {
	err := j.writer.Encode(od)
	if err != nil {
		printErr(od.Source, err)
	}
}
