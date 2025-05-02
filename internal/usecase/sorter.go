package usecase

import (
	"ball-sorting/internal/delivery"
	"ball-sorting/internal/domain"
	"ball-sorting/internal/repository"
	"sort"
)

type Sorter struct {
	inputReader  repository.InputReader
	outputWriter delivery.OutputWriter
}

func NewSorter(inputReader repository.InputReader, outputWriter delivery.OutputWriter) *Sorter {
	return &Sorter{
		inputReader:  inputReader,
		outputWriter: outputWriter,
	}
}

func (s *Sorter) Run() error {
	input, err := s.inputReader.ReadInput()
	if err != nil {
		s.outputWriter.WriteError(err)
		return err
	}

	result := &domain.Output{
		CanSort: s.canSort(input),
	}

	s.outputWriter.WriteResult(result)
	return nil
}

func (s *Sorter) canSort(input *domain.Input) bool {
	n := input.N
	containerSums := make([]int, n)
	colorSums := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := input.Containers[i][j]
			containerSums[i] += val
			colorSums[j] += val
		}
	}

	sort.Ints(containerSums)
	sort.Ints(colorSums)

	for i := 0; i < n; i++ {
		if containerSums[i] != colorSums[i] {
			return false
		}
	}

	return true
}
