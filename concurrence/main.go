package main

import (
	"fmt"
	"sync"
	"time"
)

// Structure pour représenter une opération
type Operation struct {
	ID     int
	Nom    string
	Duree  int
}

func executerOperation(op Operation, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Opération %d (%s) commencée\n", op.ID, op.Nom)
	time.Sleep(time.Duration(op.Duree) * time.Second)

	resultat := fmt.Sprintf("Opération %d (%s) terminée en %ds", op.ID, op.Nom, op.Duree)
	ch <- resultat
}

func main() {
	// Créer un channel
	resultats := make(chan string)

	// WaitGroup pour synchroniser les goroutines
	var wg sync.WaitGroup

	// Liste des opérations
	operations := []Operation{
		{ID: 1, Nom: "Téléchargement", Duree: 3},
		{ID: 2, Nom: "Traitement", Duree: 2},
		{ID: 3, Nom: "Validation", Duree: 4},
		{ID: 4, Nom: "Sauvegarde", Duree: 1},
		{ID: 5, Nom: "Synchronisation", Duree: 2},
	}

	// Lancer les goroutines
	for _, op := range operations {
		wg.Add(1)
		go executerOperation(op, resultats, &wg)
	}

	// Goroutine pour fermer le channel quand tout est fini
	go func() {
		wg.Wait()
		close(resultats)
	}()

	// Afficher tous les résultats
	fmt.Println("========== EN COURS ==========")
	for resultat := range resultats {
		fmt.Println(resultat)
	}

	fmt.Println("=============================")
	fmt.Println("✨ Toutes les opérations sont terminées!")
}
