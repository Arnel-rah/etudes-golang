package main

import (
    "fmt"
    "time"
)

func executerTache(id int, ch chan string) {
    fmt.Printf("Tâche %d commencée\n", id)

    time.Sleep(time.Duration(2+id) * time.Second)

    // TODO: Envoyer le résultat dans le channel

    fmt.Printf("Tâche %d terminée\n", id)
}

func main() {
    // TODO: Créer un channel

    // TODO: Lancer 5 goroutines

    // TODO: Recevoir les résultats

    // TODO: Afficher le résumé
}
