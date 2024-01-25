package main

import (
	"container/heap"
	"fmt"
	"math"
)

type node struct {
	index int
	dist  int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func dijkstra(graph [][]node, start int, end int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		visited[i] = false
	}

	dist[start] = 0

	pq := make(priorityQueue, 0)
	heap.Push(&pq, &node{start, 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*node)

		if visited[u.index] {
			continue
		}

		visited[u.index] = true

		for _, v := range graph[u.index] {
			if !visited[v.index] && dist[u.index]+v.dist < dist[v.index] {
				dist[v.index] = dist[u.index] + v.dist
				heap.Push(&pq, &node{v.index, dist[v.index]})
			}
		}
	}

	return dist
}

func main() {
	graph := [][]node{
		{{1, 7}, {2, 5}, {3, 2}},
		{{0, 7}, {4, 3}},
		{{0, 5}, {3, 10}, {4, 4}},
		{{0, 2}, {2, 10}, {5, 2}},
		{{1, 3}, {2, 4}, {5, 6}},
		{{1, 8}, {3, 2}, {4, 6}},
	}

	fmt.Println("The given nodes are:", graph)
	start := 2
	end := 3
	dist := dijkstra(graph, start, end)
	fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
}
