package domain

// Input - входные данные задачи
type Input struct {
	N          int       // Количество контейнеров/цветов
	Containers [][]int   // Матрица: Containers[контейнер][цвет] = количество
}

// Output - результат проверки
type Output struct {
	CanSort bool
}