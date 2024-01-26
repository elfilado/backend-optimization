package main

import (
	"fmt"
	"math"
)

type nodeX struct {
	dest   int
	weight int
}

func ConvertToDijkstraRepresentation(gameGrid [][]int) [][]nodeX {
	rows := len(gameGrid)
	cols := len(gameGrid[0])
	numVertices := rows * cols

	dijkstraMatrix := make([][]nodeX, numVertices)

	for i := 0; i < numVertices; i++ {
		dijkstraMatrix[i] = make([]nodeX, 0)
	}

	// Helper function to add an edge to the Dijkstra matrix
	addEdge := func(src, dest, weight int) {
		dijkstraMatrix[src] = append(dijkstraMatrix[src], nodeX{dest: dest, weight: weight})
		dijkstraMatrix[dest] = append(dijkstraMatrix[dest], nodeX{dest: src, weight: weight})
	}

	// Iterate through the game grid to create edges in the Dijkstra matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check if the current cell is not empty (0)
			if gameGrid[i][j] != 0 {
				// Check neighboring cells to add edges
				for row := -1; row <= 1; row++ {
					for col := -1; col <= 1; col++ {
						newRow, newCol := i+row, j+col
						// Check if the neighbor is within bounds and not the current cell
						if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && (newRow != i || newCol != j) {
							// Check if the neighbor is not empty (0) and add an edge
							if gameGrid[newRow][newCol] != 0 {
								src := i*cols + j
								dest := newRow*cols + newCol
								weight := int(math.Abs(float64(gameGrid[i][j] - gameGrid[newRow][newCol])))
								addEdge(src, dest, weight)
							}
						}
					}
				}
			}
		}
	}

	// Convert the dijkstraMatrix to the desired representation
	dijkstraRepresentation := make([][]nodeX, numVertices)
	for i, neighbors := range dijkstraMatrix {
		dijkstraRepresentation[i] = neighbors
	}

	return dijkstraRepresentation
}

func main() {
	/*gameGrid := [][]int{
		{0, 1, 0, 2},
		{1, 0, 2, 0},
		{0, 2, 0, 1},
		{2, 0, 1, 0},
	}*/
	gameGrid := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	dijkstraRepresentation := ConvertToDijkstraRepresentation(gameGrid)

	// Print the resulting Dijkstra representation
	for i, neighbors := range dijkstraRepresentation {
		fmt.Printf("%d: %v\n", i, neighbors)
	}
}
