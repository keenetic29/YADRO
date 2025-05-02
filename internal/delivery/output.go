package delivery

import (
	"ball-sorting/internal/domain"
	"os"
)

type OutputWriter interface {
	WriteResult(result *domain.Output)
	WriteError(err error)
}

type StdoutOutput struct{}

func NewStdoutOutput() OutputWriter {
	return &StdoutOutput{}
}

func (o *StdoutOutput) WriteResult(result *domain.Output) {
	if result.CanSort {
		os.Stdout.WriteString("yes\n")
	} else {
		os.Stdout.WriteString("no\n")
	}
}

func (o *StdoutOutput) WriteError(err error) {
	os.Stderr.WriteString("Ошибка: " + err.Error() + "\n")
}