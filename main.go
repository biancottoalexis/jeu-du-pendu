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
		fmt.Println("Erreur : la reprise n'est pas encore faite.")
		os.Exit(1)
	} else if commande == "score" {
		fmt.Println("Erreur : les scores ne sont pas encore faits.")
		os.Exit(1)
	} else {
		fmt.Println("Usage : go run . <jouer|reprendre|score>")
		os.Exit(1)
	}
}