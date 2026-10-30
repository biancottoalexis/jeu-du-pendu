package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Sauvegarde struct {
	Pseudo  string   `json:"pseudo"`
	Mot     string   `json:"mot"`
	Lettres []string `json:"lettres"`
	Essais  int      `json:"essais"`
}

func sauvegarder(pseudo string, mot string, lettres string, essais int) {
	save := Sauvegarde{
		Pseudo:  pseudo,
		Mot:     mot,
		Lettres: strings.Split(lettres, ""),
		Essais:  essais,
	}

	contenu, err := json.Marshal(save)
	if err != nil {
		fmt.Println("Erreur")
		return
	}

	err = os.WriteFile("save.json", data, 0644)
	if err != nil {
		fmt.Println("Erreur")
		return
	}

	fmt.Println("Partie sauvegardée")
}

func reprendre() {
	data, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("Aucune sauvegarde")
		return
	}

	var save Sauvegarde

	err = json.Unmarshal(data, &save)
	if err != nil {
		fmt.Println("Erreur")
		return
	}

	os.Remove("save.json")

	letters := strings.Join(save.Lettres, "")

	IndicateLetter(
		save.Pseudo,
		save.Mot,
		letters,
		save.Essais,
	)
}
