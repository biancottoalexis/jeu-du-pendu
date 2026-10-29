package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func choisirMot(nomFichier string) string {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur de lecture de", nomFichier)
		return ""
	}

	mots := strings.Fields(string(contenu))
	if len(mots) == 0 {
		fmt.Println("Le fichier", nomFichier, "est vide.")
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return mots[r.Intn(len(mots))]
}

func selectionanimaux() string {
	return choisirMot("mot-animaux.txt")
}

func selectionmetiers() string {
	return choisirMot("mot-metiers.txt")
}

func selectionpays() string {
	return choisirMot("mot-pays.txt")
}

func selectionnourriture() string {
	return choisirMot("mot-nourriture.txt")
}

func selectionsimple() string {
	return choisirMot("mot-simple.txt")
}

func selectionhard() string {
	return choisirMot("mot-hard.txt")
}
