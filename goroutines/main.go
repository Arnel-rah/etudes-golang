package main

import (
	"fmt"
	"time"
)

func executerTache(id int, ch chan string) {
	fmt.Printf("Tâche %d commencée\n", id)
	time.Sleep(time.Duration(2+id) * time.Second)

	// Envoyer le résultat dans le channel
	resultat := fmt.Sprintf("Tâche %d terminée avec succès", id)
	ch <- resultat

	fmt.Printf("Tâche %d terminée\n", id)
}

func main() {
	// Créer un channel
	resultats := make(chan string)

	// Lancer 5 goroutines
	for i := 1; i <= 5; i++ {
		go executerTache(i, resultats)
	}

	// Recevoir les résultats
	for i := 0; i < 5; i++ {
		resultat := <-resultats
		fmt.Println(resultat)
	}

	// Afficher le résumé
	fmt.Println("=============================")
	fmt.Println("Résumé : 5 tâches complétées avec succès!")
}
