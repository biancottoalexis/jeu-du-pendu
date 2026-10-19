package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func selection() {
	contenu, err := os.ReadFile("mot.txt")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	mots := strings.Split(strings.TrimSpace(string(contenu)), "\n")
	mot := mots[rand.Intn(len(mots))]

	fmt.Println(mot)
}
