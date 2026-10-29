package main

import (
	"fmt"
	"strings"
)

func newgames() {
	fmt.Println("===================================")
	fmt.Println("   BIENVENUE DANS LE JEU DU PENDU")
	fmt.Println("===================================")
	fmt.Println()

	fmt.Print("Voulez-vous lire les regles ? (oui / non) : ")

	var choix string
	fmt.Scanln(&choix)
	choix = strings.ToLower(choix)

	if choix == "oui" {
		fmt.Println()
		fmt.Println("========== REGLES DU JEU ==========")
		fmt.Println(`
- Le programme choisit un mot au hasard.
- Le joueur propose une lettre a chaque tour.
- Si la lettre est dans le mot, ses occurrences sont revelees.
- Une mauvaise lettre fait perdre 1 essai.
- Une entree invalide fait perdre 2 essais.
- Une lettre deja proposee ne fait perdre aucun essai.
- Le joueur dispose de 6 essais.
- Les majuscules et minuscules sont equivalentes.
- Taper !save sauvegarde la partie et quitte le jeu.
- Le jeu utilise mot.txt, save.json et scores.csv.
`)
	} else if choix == "non" {
		fmt.Println("D'accord, on passe les regles !")
	} else {
		fmt.Println("Choix invalide. Entrez oui ou non.")
		return
	}

	fmt.Println()
	fmt.Println("========== NOUVELLE PARTIE ==========")
	fmt.Println()

	var pseudo string
	fmt.Print("Entrez votre pseudo : ")
	fmt.Scanln(&pseudo)

	if pseudo == "" {
		fmt.Println("Erreur : entrez un pseudo valide.")
		return
	}

	fmt.Println()
	fmt.Println("Bienvenue", pseudo, "!")
	fmt.Println()

	fmt.Println("Choisissez la categorie :")
	fmt.Println("1. Animaux")
	fmt.Println("2. Metiers")
	fmt.Println("3. Pays")
	fmt.Println("4. Nourriture")
	fmt.Println("5. Simple")
	fmt.Println("6. Hard")
	fmt.Print("Entrez 1, 2, 3, 4, 5 ou 6 : ")

	fmt.Scanln(&choix)

	var mot string
	var nomCategorie string

	if choix == "1" {
		nomCategorie = "Animaux"
		mot = selectionanimaux()
	} else if choix == "2" {
		nomCategorie = "Metiers"
		mot = selectionmetiers()
	} else if choix == "3" {
		nomCategorie = "Pays"
		mot = selectionpays()
	} else if choix == "4" {
		nomCategorie = "Nourriture"
		mot = selectionnourriture()
	} else if choix == "5" {
		nomCategorie = "Simple"
		mot = selectionsimple()
	} else if choix == "6" {
		nomCategorie = "Hard"
		mot = selectionhard()
	} else {
		fmt.Println("Choix invalide.")
		return
	}

	if mot == "" {
		fmt.Println("Impossible de charger un mot.")
		return
	}

	fmt.Println()
	fmt.Println("Nouvelle partie !")
	fmt.Println("Categorie choisie :", nomCategorie)
	fmt.Println("Le mot contient", len([]rune(mot)), "lettres.")
    fmt.Println("Vous avez 6 essais pour deviner le mot.")
    fmt.Println("taper !save pour sauvegarder la partie et quitter le jeu.")
	fmt.Println("Proposez une lettre.")

	IndicateLetter(mot)
	if saisie == "!save" {
		fmt.Println("Partie sauvegardée. A bientot !")
		return
	}

}
