package main

import (
	"fmt"
	"math" // For using math.MaxInt32
)

// Function that take in parameters graph, start node and end node
// And returns the shortest distance from start node to end node
func dijkstra(graph [][]int, start int, end int) []int {
	n := len(graph)            // Number of nodes in the graph
	dist := make([]int, n)     // Array to store the shortest distance from start node to all other nodes
	visited := make([]bool, n) // Array to keep track of visited nodes

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		visited[i] = false
	}

	dist[start] = 0 // Distance from start node to itself is 0

	for count := 0; count < n-1; count++ {
		u := -1 // Initialize u to -1

		for i := 0; i < n; i++ { // Find the node with the shortest distance from the set of nodes not yet visited
			if !visited[i] && (u == -1 || dist[i] < dist[u]) {
				u = i
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true

		for v := 0; v < n; v++ { // Update the distance of the nodes adjacent to u
			if graph[u][v] != 0 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v] // Update the distance of node v
			}
		}
	}

	return dist // Return the shortest distance from start node to all other nodes
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

	graph := [][]int{
		{0, 7, 5, 2, 0, 0, 2, 4, 3},
		{7, 0, 0, 0, 3, 0, 0, 0, 0},
		{5, 0, 0, 10, 4, 0, 0, 0, 0},
		{2, 0, 10, 0, 0, 2, 0, 0, 0},
		{0, 3, 4, 0, 0, 6, 0, 0, 0},
		{0, 8, 0, 2, 6, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 5, 0},
		{4, 0, 0, 0, 0, 0, 5, 0, 1},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
	}

	fmt.Println("The given nodes are:", graph)
	start := 2
	end := 3

	dist := dijkstra(graph, start, end)

	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
}
