package hexgames

import (
	"container/heap"
	"math"
	"runtime"
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

func dijkstra(graph [][]node, start int) []int {
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

func GridToGraph(grid [][]int) [][]node {
	longueur := len(grid)
	hauteur := len(grid[0])

	graph := [][]node{}

	for y := range grid {
		for x := range grid[y] {
			var nodeD []node

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					// remove la diagonale de en haut à gauche à en bas à droite
					if (dx == -1 && dy == -1) || (dx == 1 && dy == 1 || (dx == 0 && dy == 0)) {
						continue
					}
					// Supprime les cases en dehors du tableau
					if x+dx < 0 || x+dx >= longueur || y+dy < 0 || y+dy >= hauteur {
						continue
					}
					// Supprime les cases avec des enemis
					if grid[y+dy][x+dx] == 2 {
						continue
					}

					dist := 1
					// Si c'est un jeton a nous on met la distant a zero
					if grid[y+dy][x+dx] == 1 {
						dist = 0
					}

					nodeD = append(nodeD, node{(y+dy)*longueur + x + dx, dist})
				}
			}

			graph = append(graph, nodeD)
		}
	}

	return graph
}

func GridToGraphOptiIf(grid [][]int) [][]node {
	longueur := len(grid)
	hauteur := len(grid[0])

	graph := [][]node{}

	adjList := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{0, -1},
		{-1, -1},
		{-1, 0},
	}

	for y := range grid {
		for x := range grid[y] {
			var nodeD []node

			for _, adj := range adjList {
				// Supprime les cases en dehors du tableau
				if x+adj[0] < 0 || x+adj[0] >= longueur || y+adj[1] < 0 || y+adj[1] >= hauteur {
					continue
				}
				// Supprime les cases avec des enemis
				if grid[y+adj[1]][x+adj[0]] == 2 {
					continue
				}

				dist := 1
				// Si c'est un jeton a nous on met la distant a zero
				if grid[y+adj[1]][x+adj[0]] == 1 {
					dist = 0
				}

				nodeD = append(nodeD, node{(y+adj[1])*longueur + x + adj[0], dist})
			}

			graph = append(graph, nodeD)
		}
	}

	return graph
}

func GridToMinDist(grid [][]int) int {
	graph := GridToGraph(grid)

	longueur := len(grid)
	hauteur := len(grid[0])

	minDist := math.MaxInt32

	//for sur la première ligne et la dernière ligne du tableau
	for y := range grid[0] {
		for x := range grid[len(grid)-1] {
			start := y
			end := longueur*(hauteur-1) + x

			data := dijkstra(graph, start)
			if data[end] < minDist {
				minDist = data[end]
			}
		}
	}

	return minDist
}

func GridToMinDistWithGoRoutine(grid [][]int) int {
	longueur := len(grid)
	hauteur := len(grid[0])

	graph := GridToGraph(grid)

	minDist := math.MaxInt32

	// Design Pattern : Sémaphore
	// Traiter de manière la plus optimisée possible les go routines pour éviter les pools de go routines

	// Runtime.NumCPU récupère le nombre de coeurs sur la machine
	maxWorkers := runtime.NumCPU()
	// On le passe comme indicateur de taille dans un channel de struct. Une struct vide pour le moins de mémoire possible (plus efficace).
	// Le nombre de coeurs détermine l'optimisation en termes du nombre de lancements simultanés de go routines (cycles processeur).
	sem := make(chan struct{}, maxWorkers)
	results := make(chan int, 100) // canal pour les résultats

	// Système de jetons
	//for sur la première ligne et la dernière ligne du tableau
	for y := range grid[0] {
		for x := range grid[len(grid)-1] {
			sem <- struct{}{} // Acquiert un jeton
			x := x
			go func() {
				defer func() { <-sem }() // Libère un jeton
				start := y
				end := longueur*(hauteur-1) + x

				ff := dijkstra(graph, start)
				results <- ff[end]
			}()
		}
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}

	close(sem)
	close(results)

	for result := range results {
		if result < minDist {
			minDist = result
		}
	}

	return minDist
}

func GridToMinDistWithGoRoutineOptimisation(grid [][]int) int {
	graph := GridToGraph(grid)

	longueur := len(grid)
	hauteur := len(grid[0])

	minDist := math.MaxInt32

	// Design Pattern : Sémaphore
	// Traiter de manière la plus optimisée possible les go routines pour éviter les pools de go routines

	// Runtime.NumCPU récupère le nombre de coeurs sur la machine
	maxWorkers := runtime.NumCPU()
	// On le passe comme indicateur de taille dans un channel de struct. Une struct vide pour le moins de mémoire possible (plus efficace).
	// Le nombre de coeurs détermine l'optimisation en termes du nombre de lancements simultanés de go routines (cycles processeur).
	sem := make(chan struct{}, maxWorkers)
	results := make(chan int, 100) // canal pour les résultats

	// Système de jetons
	//for sur la première ligne et la dernière ligne du tableau
	for y := range grid[0] {
		sem <- struct{}{} // Acquiert un jeton
		go func(start int) {
			defer func() { <-sem }() // Libère un jeton

			ppp := dijkstra(graph, start)

			for x := range grid[len(grid)-1] {
				end := longueur*(hauteur-1) + x
				//fmt.Println(ppp)
				/*dist := */
				results <- ppp[end]
			}
		}(y)
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}

	close(sem)
	close(results)

	for result := range results {
		if result < minDist {
			minDist = result
		}
	}

	return minDist
}
