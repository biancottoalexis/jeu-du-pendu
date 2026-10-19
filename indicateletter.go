package main

import (
	"fmt"
)

func IndicateLetter(mot string) {
	essaie := 6
	trouver := false
	for essaie > 0 && !trouver {
		fmt.Print("Enter a letter: ")
		var lettre string
		fmt.Scan(&lettre)
	}
}
