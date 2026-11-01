package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

func ajouterScore(pseudo string, statut string, essais int, mot string, lettres string) {
	fichier, err := os.OpenFile("scores.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur")
		return
	}

	info, err := fichier.Stat()
	if err != nil {
		fmt.Println("Erreur")
		fichier.Close()
		return
	}

	ecrire := csv.NewWriter(fichier)

	if info.Size() == 0 {
		ecrire.Write([]string{"pseudo", "statut", "essais_restants", "mot", "lettres", "date"})
	}

	date := time.Now().Format("2006-01-02 15:04")

	score := []string{pseudo, statut, fmt.Sprint(essais), mot, lettres, date}

	ecrire.Write(score)
	ecrire.Flush()

	if ecrire.Error() != nil {
		fmt.Println("Erreur")
	}

	fichier.Close()
}

func afficherScores() {
	fichier, err := os.Open("scores.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Aucun score")
			return
		}

		fmt.Println("Erreur")
		return
	}

	lire := csv.NewReader(fichier)

	fmt.Println("PSEUDO\tSTATUT\tESSAIS\tMOT\tDATE")

	numero := 0

	for {
		ligne, err := lire.Read()

		if err == io.EOF {
			break
		}

		numero++

		if err != nil {
			fmt.Println("Ligne", numero, "ignorée")
			continue
		}

		if len(ligne) != 6 {
			fmt.Println("Ligne", numero, "ignorée")
			continue
		}

		if ligne[0] == "pseudo" {
			continue
		}

		fmt.Println(ligne[0], "\t", ligne[1], "\t", ligne[2], "\t", ligne[3], "\t", ligne[5])
	}

	fichier.Close()
}

// guidé par l'ia
