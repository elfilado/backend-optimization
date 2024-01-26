package main

import (
	"fmt"
	"runtime"
	"time"
)

// Design Pattern : Sémaphore
// Traiter de manière la plus optimisée possible les go routines pour éviter les pools de go routines
func semaphore() {
	// Runtime.NumCPU récupère le nombre de coeurs sur la machine
	maxWorkers := runtime.NumCPU()
	// On le passe comme indicateur de taille dans un channel de struct. Une struct vide pour le moins de mémoire possible (plus efficace).
	// Le nombre de coeurs détermine l'optimisation en termes du nombre de lancements simultanés de go routines (cycles processeur).
	sem := make(chan struct{}, maxWorkers)
	results := make(chan string, 100) // canal pour les résultats

	// Système de jetons
	for i := 0; i < 100; i++ {
		sem <- struct{}{} // Acquiert un jeton
		go func(id int) {
			defer func() { <-sem }() // Libère un jeton
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Goroutine ", id, " done : sem len = ", len(sem))
		}(i)
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}

	close(sem)
}
