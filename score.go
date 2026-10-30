package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func ajouterScore(pseudo string, statut string, essais int, mot string, lettres string) {
	fichier, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur")
		return
	}
	defer fichier.Close()

	ecrire := csv.NewWriter(fichier)

	date := time.Now().Format("2006-01-02 15:04")

	score := []string{
		pseudo,
		statut,
		fmt.Sprint(essais),
		mot,
		lettres,
		date,
	}

	ecrire.Write(score)
	ecrire.Flush()
}

func afficherScores() {
	fichier, err := os.Open("scores.csv")
	if err != nil {
		fmt.Println("Aucun score")
		return
	}
	defer fichier.Close()

	lire := csv.NewReader(fichier)

	lignes, err := lire.ReadAll()
	if err != nil {
		fmt.Println("Erreur")
		return
	}

	fmt.Println("PSEUDO\tSTATUT\tESSAIS\tMOT\tDATE")

	for _, ligne := range lignes {
		if len(ligne) < 6 {
			continue
		}

		fmt.Println(
			ligne[0], "\t",
			ligne[1], "\t",
			ligne[2], "\t",
			ligne[3], "\t",
			ligne[5],
		)
	}
}

// aide à l'intelligence artificiel et documentation golang 