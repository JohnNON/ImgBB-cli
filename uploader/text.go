package uploader

import (
	"fmt"
	"os"
)

type textWriter struct{}

func newTextWriter() *textWriter {
	return &textWriter{}
}

func (t *textWriter) write(od outData) {
	fmt.Fprintf(os.Stdout, "%-v\n", od)
}
