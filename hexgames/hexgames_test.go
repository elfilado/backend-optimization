package hexgames

import "testing"

var grid = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 2, 0, 0, 0, 0},
	{0, 2, 0, 1, 2, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 2, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 0, 2, 0},
	{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 2, 0, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func BenchmarkGridToGraphOptiIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GridToGraphOptiIf(grid)
	}
}

func BenchmarkGridToGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GridToGraph(grid)
	}
}

func BenchmarkGridToMinDist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GridToMinDist(grid)
	}
}

func BenchmarkGridToMinDistWithGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GridToMinDistWithGoRoutine(grid)
	}
}

func BenchmarkGridToMinDistWithGoRoutineOptimisation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GridToMinDistWithGoRoutineOptimisation(grid)
	}
}
