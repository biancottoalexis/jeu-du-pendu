package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func choisirMot(nomFichier string) string {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier")
		return ""
	}

	lignes := strings.Split(string(contenu), "\n")

	var mots []string

	for _, mot := range lignes {
		mot = strings.TrimSpace(mot)
		mot = strings.ToLower(mot)

		if mot == "" {
			continue
		}

		if !estValide(mot) {
			continue
		}

		mots = append(mots, mot)
	}

	if len(mots) == 0 {
		fmt.Println("Aucun mot valide")
		return ""
	}

	index := rand.IntN(len(mots))

	return mots[index]
}

func estValide(mot string) bool {
	for _, lettre := range mot {
		if (lettre < 'a' || lettre > 'z') && lettre != '-' {
			return false
		}
	}

	return true
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
