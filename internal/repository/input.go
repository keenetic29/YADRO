package repository

import (
	"ball-sorting/internal/domain"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputReader interface {
	ReadInput() (*domain.Input, error)
}

type StdinInput struct{}

func NewStdinInput() InputReader {
	return &StdinInput{}
}

func (r *StdinInput) ReadInput() (*domain.Input, error) {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение N
	if !scanner.Scan() {
		return nil, fmt.Errorf("ошибка чтения N")
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 {
		return nil, fmt.Errorf("некорректное значение N")
	}

	// Чтение контейнеров
	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("ошибка чтения контейнера %d", i+1)
		}
		values := strings.Fields(scanner.Text())
		if len(values) != n {
			return nil, fmt.Errorf("неверное количество значений в контейнере %d", i+1)
		}

		container := make([]int, n)
		for j, val := range values {
			count, err := strconv.Atoi(val)
			if err != nil || count < 0 {
				return nil, fmt.Errorf("некорректное количество шаров в контейнере %d", i+1)
			}
			container[j] = count
		}
		containers[i] = container
	}

	return &domain.Input{
		N:          n,
		Containers: containers,
	}, nil
}