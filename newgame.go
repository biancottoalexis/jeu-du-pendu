package main

import (
	"fmt"
	"strings"
)

func newgames() {
	for {
		fmt.Println("===================================")
		fmt.Println("BIENVENUE DANS LE JEU DU PENDU")
		fmt.Println("===================================")
		fmt.Println()
		fmt.Println("1. Nouvelle partie")
		fmt.Println("2. Reprendre une partie")
		fmt.Println("3. Quitter")
		fmt.Print("Votre choix : ")

		var choix string
		fmt.Scanln(&choix)

		if choix == "1" {
			nouvellePartie()
		} else if choix == "2" {
			fmt.Println("La reprise de partie sera ajoutee plus tard.")
		} else if choix == "3" {
			fmt.Println("Au revoir !")
			return
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}

func nouvellePartie() {
	fmt.Println()
	fmt.Print("Voulez-vous lire les regles ? (oui / non) : ")

	var choix string

	for {
		fmt.Scanln(&choix)
		choix = strings.ToLower(choix)

		if choix == "oui" {
			fmt.Println()
			fmt.Println("========== REGLES DU JEU ==========")
			fmt.Print(`
- Le programme choisit un mot au hasard.
- Le joueur propose une lettre a chaque tour.
- Si la lettre est dans le mot, ses occurrences sont revelees.
- Une mauvaise lettre fait perdre 1 essai.
- Une entree invalide fait perdre 2 essais.
- Une lettre deja proposee ne fait perdre aucun essai.
- Le joueur dispose de 6 essais.
- Les majuscules et minuscules sont equivalentes.
- Taper !save sauvegarde la partie.
- Taper !menu revient au menu principal.
`)
			break
		} else if choix == "non" {
			fmt.Println("D'accord, on passe les regles !")
			break
		} else {
			fmt.Println("Choix invalide. Entrez oui ou non.")
			fmt.Print("Voulez-vous lire les regles ? (oui / non) : ")
		}
	}

	fmt.Println()
	fmt.Println("========== NOUVELLE PARTIE ==========")
	fmt.Println()

	var pseudo string

	for pseudo == "" {
		fmt.Print("Entrez votre pseudo : ")
		fmt.Scanln(&pseudo)

		if pseudo == "" {
			fmt.Println("Erreur : entrez un pseudo valide.")
		}
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

	var mot string
	var nomCategorie string

	for {
		fmt.Print("Entrez 1, 2, 3, 4, 5 ou 6 : ")
		fmt.Scanln(&choix)

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
			continue
		}

		break
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
	fmt.Println("Tapez !save pour sauvegarder la partie.")
	fmt.Println("Tapez !menu pour revenir au menu principal.")
	fmt.Println("Proposez une lettre.")

	IndicateLetter(pseudo, mot, lettreDepart(mot), 6)

	fmt.Println()
	fmt.Println("========== FIN DE LA PARTIE ==========")
	fmt.Println("1. Rejouer le meme mot")
	fmt.Println("2. Faire une nouvelle partie")
	fmt.Println("3. Retourner au menu principal")

	fmt.Print("Votre choix : ")

	var fin string
	fmt.Scanln(&fin)

	if fin == "1" {
		IndicateLetter(pseudo, mot, lettreDepart(mot), 6)
	} else if fin == "2" {
		nouvellePartie()
	} else if fin == "3" {
		return
	} else {
		fmt.Println("Choix invalide.")
	}
}
