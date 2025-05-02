package domain

type Input struct {
	N          int       
	Containers [][]int   
}

type Output struct {
	CanSort bool
}