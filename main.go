// The main file for the Hex Games project.

package main

import (
	"fmt"

	"github.com/elfilado/backend-optimization/hexgames"
)

func main() {
	// 0 = vide
	// 1 = nous
	// 2 = enemis
	gird := [][]int{
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{0, 0, 1, 0},
	}

	//grid := make([][]int, 1_000_000)

	fmt.Println(hexgames.GridToMinDistWithGoRoutineOptimisation(gird))
}
