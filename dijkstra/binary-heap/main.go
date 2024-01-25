package main

import (
	"container/heap"
	"fmt"
)

const inf = int(^uint(0) >> 1)

type Graph [][]int

func dijkstra(graph Graph, start, end int) ([]int, []int) {
	n := len(graph)
	dist := make([]int, n)
	prev := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = inf
		prev[i] = -1
		visited[i] = false
	}

	dist[start] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{start, 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Node)

		if visited[u.ID] {
			continue
		}
		visited[u.ID] = true

		for v := 0; v < n; v++ {
			if weight := graph[u.ID][v]; weight > 0 && dist[u.ID]+weight < dist[v] {
				dist[v] = dist[u.ID] + weight
				prev[v] = u.ID
				heap.Push(&pq, &Node{v, dist[v]})
			}
		}
	}

	// Reconstruct the path from end to start
	path := []int{}
	current := end
	for current != -1 {
		path = append([]int{current}, path...)
		current = prev[current]
	}

	return dist, path
}

func main() {
	graph := Graph{
		{0, 7, 5, 2, 0, 0},
		{7, 0, 0, 0, 3, 0},
		{5, 0, 0, 10, 4, 0},
		{2, 0, 10, 0, 0, 2},
		{0, 3, 4, 0, 0, 6},
		{0, 8, 0, 2, 6, 0},
	}

	fmt.Println("The given nodes are:", graph)
	start := 2
	end := 3

	distances, path := dijkstra(graph, start, end)

	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, distances[end])
	fmt.Printf("Path: %v\n", path)
}

// Node represents a node in the graph with an ID and distance value.
type Node struct {
	ID       int
	Distance int
}

// PriorityQueue is a priority queue implementation for Dijkstra's algorithm.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
