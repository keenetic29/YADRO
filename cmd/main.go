package main

import (
	"ball-sorting/internal/delivery"
	"ball-sorting/internal/repository"
	"ball-sorting/internal/usecase"
	"os"
)

func main() {
	// Инициализация зависимостей
	inputReader := repository.NewStdinInput()
	outputWriter := delivery.NewStdoutOutput()
	sorter := usecase.NewSorter(inputReader, outputWriter)

	// Запуск
	if err := sorter.Run(); err != nil {
		os.Exit(1)
	}
}