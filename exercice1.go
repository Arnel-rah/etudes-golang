package main

import "fmt"

func main() {
    var metier string = "Backend Engineer" 
    annee := 2026 
	teste := true

	if teste {
		fmt.Println("Test réussi !")
	}

    fmt.Printf("Objectif : %s en %d\n", metier, annee)
}