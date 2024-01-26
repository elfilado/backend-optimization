package main

import (
	"fmt"
	"time"
)

package main

import (
"fmt"
"time"
)

func maGoroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine :", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("Début du programme")

	// Lancer une goroutine
	go maGoroutine()

	// Continuer l'exécution dans la fonction main
	for i := 0; i < 3; i++ {
		fmt.Println("Fonction main :", i)
		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println("Fin du programme")
}
/* Sortie terminal :
Début du programme
Fonction main : 0
Goroutine : 0
Fonction main : 1
Goroutine : 1
Fonction main : 2
Fin du programme*/

