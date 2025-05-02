package delivery

import (
	"ball-sorting/internal/domain"
	"os"
)

type StdoutOutput struct{}

func NewStdoutOutput() domain.OutputWriter {
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