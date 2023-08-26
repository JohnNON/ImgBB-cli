package uploader

import (
	"os"

	"gopkg.in/yaml.v3"
)

type yamlWriter struct {
	writer *yaml.Encoder
}

func newYamlWriter() *yamlWriter {
	return &yamlWriter{
		writer: yaml.NewEncoder(os.Stdout),
	}
}

func (y *yamlWriter) write(od outData) {
	err := y.writer.Encode(od)
	if err != nil {
		printErr(od.Source, err)
	}
}
