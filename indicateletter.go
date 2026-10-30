package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func IndicateLetter(pseudo string, mot string, proposees string, essaie int) bool {
	
	scanner := bufio.NewScanner(os.Stdin)

	for essaie > 0 {
		masque := ""

		for _, c := range mot {
			lettre := string(c)

			if lettre == "-" || strings.Contains(proposees, lettre) {
				masque += lettre + " "
			} else {
				masque += "_ "
			}
		}

		if !strings.Contains(masque, "_") {
			fmt.Println()
			fmt.Println("Bravo, le mot était", mot, "!")
			return false
		}

		fmt.Println()

		// dessin fait par Ia en bas
		largeur := len([]rune("Mot :"+masque)) + 8
		if largeur < 30 {
			largeur = 30
		}

		afficher := func(gauche, droite string) {
			fmt.Printf("%-*s%s\n", largeur, gauche, droite)
		}

		afficher("Mot : "+masque, "+---+")
		afficher(fmt.Sprintf("Essais restants : %d", essaie), "|   |")

		if essaie <= 5 {
			afficher("", "|   O")
		} else {
			afficher("", "|")
		}

		if essaie <= 2 {
			afficher("", "|  /|\\")
		} else if essaie <= 3 {
			afficher("", "|  /|")
		} else if essaie <= 4 {
			afficher("", "|   |")
		} else {
			afficher("", "|")
		}

		if essaie <= 1 {
			afficher("", "|  / \\")
		} else if essaie <= 2 {
			afficher("", "|  /")
		} else {
			afficher("", "|")
		}

		afficher("", "|")
		afficher("", "_|_")
		// fin du dessin fait par Ia

		fmt.Println()
		fmt.Println("Lettres proposées :", proposees)
		fmt.Print("> ")

		if !scanner.Scan() {
			return false
		}

		lettre := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if lettre == "" {
			fmt.Println("Erreur : entrez une lettre valide.")
			continue
		}

		if lettre == "!menu" {
			fmt.Println("Retour au menu principal...")
			return true
		}

		if lettre == "!save" {
			fmt.Println("Partie sauvegardée")
			continue
		}

		if len(lettre) > 1 {
			valide := true
			for _, c := range lettre {
				if (c < 'a' || c > 'z') && c != '-' {
					valide = false
				}
			}

			if valide == false {
				fmt.Println("Erreur : un mot ne contient que des lettres et des tirets.")
				continue
			}

			if essaie < 2 {
				fmt.Println("Erreur : il faut au moins 2 essais pour proposer un mot.")
				continue
			}

			if len(lettre) != len(mot) {
				fmt.Println("Erreur : le mot doit avoir", len(mot), "lettres.")
				continue
			}

			if lettre == mot {
				fmt.Println()
				fmt.Println("Bravo, le mot était", mot, "!")
				return false
			}

			fmt.Println("Raté, ce n'est pas le bon mot ! -2 essais")
			essaie -= 2
			continue
		}

		if len([]rune(lettre)) != 1 || lettre < "a" || lettre > "z" {
			fmt.Println("Erreur : entrez une seule lettre.")
			continue
		}

		if strings.Contains(proposees, lettre) {
			fmt.Println("Erreur : lettre déjà proposée.")
			continue
		}

		if strings.Contains(mot, lettre) {
			proposees += lettre
			fmt.Println("Bien joué !")
		} else {
			proposees += lettre
			fmt.Println("Mauvaise lettre ! -1 essai")
			essaie--
		}
	}

	fmt.Println()
	fmt.Println("Mot :", mot)
	fmt.Println()
	fmt.Println("Perdu ! Le mot était :", mot)
	fmt.Println()
	fmt.Println("       +---+")
	fmt.Println("       |   |")
	fmt.Println("       |   O")
	fmt.Println("       |  /|\\")
	fmt.Println("       |  / \\")
	fmt.Println("       |")
	fmt.Println("      _|_")
	// fin du dessin fait par Ia

	return false
}