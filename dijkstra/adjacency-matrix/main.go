package main

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, start int, end int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		visited[i] = false
	}

	dist[start] = 0

	for count := 0; count < n-1; count++ {
		u := -1

		for i := 0; i < n; i++ {
			if !visited[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true

		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}

func main() {
	// graph := [][]int{
	// 	{0, 7, 5, 2, 0, 0},
	// 	{7, 0, 0, 0, 3, 0},
	// 	{5, 0, 0, 10, 4, 0},
	// 	{2, 0, 10, 0, 0, 2},
	// 	{0, 3, 4, 0, 0, 6},
	// 	{0, 8, 0, 2, 6, 0},
	// }

	// Big Matrix : 12x12
	graph := [][]int{
		{0, 6, 7, 2, 5, 1, 7, 4, 3, 2, 6, 7},
		{6, 0, 7, 3, 7, 2, 4, 1, 5, 3, 6, 5},
		{7, 7, 0, 1, 4, 6, 3, 2, 4, 5, 4, 3},
		{2, 3, 1, 0, 4, 7, 1, 7, 4, 6, 2, 5},
		{5, 7, 4, 4, 0, 1, 2, 7, 3, 5, 3, 6},
		{1, 2, 6, 7, 1, 0, 4, 5, 3, 2, 1, 4},
		{7, 4, 3, 1, 2, 4, 0, 6, 7, 3, 5, 2},
		{4, 1, 2, 7, 7, 5, 6, 0, 3, 1, 4, 3},
		{3, 5, 4, 4, 3, 3, 7, 3, 0, 2, 4, 1},
		{2, 3, 5, 6, 5, 2, 3, 1, 2, 0, 5, 7},
		{6, 6, 4, 2, 3, 1, 5, 4, 4, 5, 0, 6},
		{7, 5, 3, 5, 6, 4, 2, 3, 1, 7, 6, 0},
	}

	fmt.Println("The given nodes are:", graph)
	start := 2
	end := 3

	dist := dijkstra(graph, start, end)

	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
}
