package main

import (
	"ball-sorting/internal/delivery"
	"ball-sorting/internal/repository"
	"ball-sorting/internal/usecase"
	"os"
)

func main() {
	inputReader := repository.NewStdinInput(os.Stdin)
	outputWriter := delivery.NewStdoutOutput(os.Stdout, os.Stderr)

	sorter := usecase.NewSorter(inputReader, outputWriter)
	if err := sorter.Run(); err != nil {
		os.Exit(1)
	}
}

