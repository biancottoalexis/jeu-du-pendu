
package main

import (
    "fmt"
)

func newgames() {
    fmt.Println("Bienvenue dans le jeu du pendu ! en go")

    pseudo := "leandro"
    fmt.Println("Bienvenue", pseudo)

    if pseudo == "" {
        fmt.Println("Erreur : Entrez un pseudo valide")
        return
    }
}
