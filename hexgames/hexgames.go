package main

import (
	"container/heap"
	"fmt"
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

func Tt() {
	gird := make([][]int, 1_000_000)

	longueur := len(gird)
	hauteur := len(gird[0])

	graph := [][]node{}

	for y := range gird {
		for x := range gird[y] {
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
					if gird[y+dy][x+dx] == 2 {
						continue
					}
					//graph[y][x] = node{(y * dy) * (x + dx), 1}
					//append

					dist := 1
					// Si c'est un jeton a nous on met la distant a zero
					if gird[y+dy][x+dx] == 1 {
						dist = 0
					}

					nodeD = append(nodeD, node{(y+dy)*longueur + x + dx, dist})
				}
			}

			/*for _, vv := range nodeD {
				fmt.Print(vv)
			}
			fmt.Println()*/

			graph = append(graph, nodeD)
		}
	}

	//fmt.Println("-----------------------")

	/*for y := range gird {
		for x := range gird[y] {
			fmt.Print(graph[y][x])
		}
		fmt.Println()
	}*/

	//for sur la première ligne et la dernière ligne du tableau
	for y := range gird[0] {
		for _ = range gird[len(gird)-1] {
			start := y
			//end := longueur*(hauteur-1) + x
			/*dist := */
			dijkstra(graph, start)
			//fmt.Printf("Shortest path from node %d to %d: %d\n", start, end, dist[end])
		}
	}
}

func T3() {
	gird := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	//gird := make([][]int, 1_000_000)

	longueur := len(gird)
	hauteur := len(gird[0])

	graph := [][]node{}

	for y := range gird {
		for x := range gird[y] {
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
					if gird[y+dy][x+dx] == 2 {
						continue
					}
					//graph[y][x] = node{(y * dy) * (x + dx), 1}
					//append

					dist := 1
					// Si c'est un jeton a nous on met la distant a zero
					if gird[y+dy][x+dx] == 1 {
						dist = 0
					}

					nodeD = append(nodeD, node{(y+dy)*longueur + x + dx, dist})
				}
			}

			/*for _, vv := range nodeD {
				fmt.Print(vv)
			}
			fmt.Println()*/

			graph = append(graph, nodeD)
		}
	}

	//fmt.Println("-----------------------")

	/*for y := range gird {
		for x := range gird[y] {
			fmt.Print(graph[y][x])
		}
		fmt.Println()
	}*/

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
	for y := range gird[0] {
		for x := range gird[len(gird)-1] {
			sem <- struct{}{} // Acquiert un jeton
			x := x
			go func() {
				defer func() { <-sem }() // Libère un jeton
				start := y
				end := longueur*(hauteur-1) + x

				ff := dijkstra(graph, start)
				fmt.Println(ff)
				/*dist := */ results <- ff[end]
			}()
		}
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}

	close(sem)
	close(results)

	for val := range results {
		fmt.Println(val)
	}
}

func T4() {
	gird := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	//gird := make([][]int, 1_000_000)

	longueur := len(gird)
	hauteur := len(gird[0])

	graph := [][]node{}

	for y := range gird {
		for x := range gird[y] {
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
					if gird[y+dy][x+dx] == 2 {
						continue
					}
					//graph[y][x] = node{(y * dy) * (x + dx), 1}
					//append

					dist := 1
					// Si c'est un jeton a nous on met la distant a zero
					if gird[y+dy][x+dx] == 1 {
						dist = 0
					}

					nodeD = append(nodeD, node{(y+dy)*longueur + x + dx, dist})
				}
			}

			/*for _, vv := range nodeD {
				fmt.Print(vv)
			}
			fmt.Println()*/

			graph = append(graph, nodeD)
		}
	}

	//fmt.Println("-----------------------")

	/*for y := range gird {
		for x := range gird[y] {
			fmt.Print(graph[y][x])
		}
		fmt.Println()
	}*/

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
	for y := range gird[0] {
		//fmt.Println(y)
		sem <- struct{}{} // Acquiert un jeton
		y := y
		go func() {
			defer func() { <-sem }() // Libère un jeton
			start := y

			//fmt.Println(start)
			ppp := dijkstra(graph, start)

			fmt.Println(ppp)

			for x := range gird[len(gird)-1] {
				end := longueur*(hauteur-1) + x
				//fmt.Println(ppp)
				/*dist := */
				results <- ppp[end]
			}
		}()
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}

	close(sem)
	close(results)
	fmt.Println("------------")
	for val := range results {
		fmt.Println(val)
	}
}

func Tt2() {
	graph := [][]node{
		{{1, 5}, {3, 9}},
		{{0, 5}, {2, 2}},
		{{1, 2}, {3, 3}, {4, 7}},
		{{0, 9}, {2, 3}},
		{{2, 7}},
	}

	//fmt.Println("The given nodes are:", graph)
	start := 0
	//end := 3
	/*dist := */
	dijkstra(graph, start)
}

func main() {
	// 0 = vide
	// 1 = nous
	// 2 = enemis
	/*gird := [][]int{
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{0, 0, 1, 0},
	}*/

	T3()
	fmt.Println("------------------------")
	fmt.Println("------------------------")
	T4()

}
