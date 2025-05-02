package usecase

import (
	"ball-sorting/internal/domain"
	"sort"
)

type Sorter struct {
	inputReader  domain.InputReader
	outputWriter domain.OutputWriter
}

func NewSorter(inputReader domain.InputReader, outputWriter domain.OutputWriter) *Sorter {
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
