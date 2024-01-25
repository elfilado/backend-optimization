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
	graph := [][]int{
		{0, 5, 0, 9, 0},
		{5, 0, 2, 0, 0},
		{0, 2, 0, 3, 7},
		{9, 0, 3, 0, 0},
		{0, 0, 7, 0, 0},
	}
	fmt.Println("The given nodes are:", graph)
	start := 0
	end := 4

	dist := dijkstra(graph, start, end)

	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
}
