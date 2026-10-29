package main

import (
"fmt"
"math/rand"
"os"
)

func choisirMot(nomFichier string) string {
contenu, err := os.ReadFile(nomFichier)
if err != nil {
fmt.Println("Erreur de lecture de", nomFichier)
return ""
}

var mots []string
mot := ""

for _, lettre := range string(contenu) {
	if lettre == ' ' || lettre == '\n' || lettre == '\t' || lettre == '\r' {
		if mot != "" {
			mots = append(mots, mot)
			mot = ""
		}
	} else {
		mot += string(lettre)
	}
}

if mot != "" {
	mots = append(mots, mot)
}

if len(mots) == 0 {
	fmt.Println("Le fichier", nomFichier, "est vide.")
	return ""
}

return mots[rand.Intn(len(mots))]

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