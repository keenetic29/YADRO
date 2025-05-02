package domain

// InputReader - интерфейс для чтения входных данных
type InputReader interface {
	ReadInput() (*Input, error)
}

// OutputWriter - интерфейс для вывода результатов
type OutputWriter interface {
	WriteResult(result *Output)
	WriteError(err error)
}

// Input - входные данные задачи
type Input struct {
	N          int       // Количество контейнеров/цветов
	Containers [][]int   // Матрица: Containers[контейнер][цвет] = количество
}

// Output - результат проверки
type Output struct {
	CanSort bool
}