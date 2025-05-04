package delivery

import (
	"ball-sorting/internal/domain"
	"io"
)

type OutputWriter interface {
	WriteResult(result *domain.Output)
	WriteError(err error)
}

type StdoutOutput struct {
	outWriter io.Writer
	errWriter io.Writer
}

func NewStdoutOutput(outWriter, errWriter io.Writer) *StdoutOutput {
	return &StdoutOutput{
		outWriter: outWriter,
		errWriter: errWriter,
	}
}

func (o *StdoutOutput) WriteResult(result *domain.Output) {
	var output string
	if result.CanSort {
		output = "yes\n"
	} else {
		output = "no\n"
	}
	o.outWriter.Write([]byte(output))
}

func (o *StdoutOutput) WriteError(err error) {
	o.errWriter.Write([]byte("Ошибка: " + err.Error() + "\n"))
}