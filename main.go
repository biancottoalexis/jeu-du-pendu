package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <jouer|reprendre|score>")
		os.Exit(1)
	}

	commande := os.Args[1]

	if commande == "jouer" {
		nouvellePartie()
	} else if commande == "reprendre" {
		reprendre()
	} else if commande == "score" {
		afficherScores()
	} else {
		fmt.Println("Usage : go run . <jouer|reprendre|score>")
		os.Exit(1)
	}
}