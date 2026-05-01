package main

import (
	"fmt"
	"sync"
	"time"
)

type CompteBancaire struct {
	solde int
	mu    sync.Mutex
}

func (c *CompteBancaire) Deposer(montant int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Printf("Dépôt de %d$ en cours...\n", montant)
	time.Sleep(500 * time.Millisecond)

	c.solde += montant
	fmt.Printf(" Dépôt réussi. Nouveau solde: %d$\n", c.solde)
}

func (c *CompteBancaire) Retirer(montant int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Printf("Retrait de %d$ en cours...\n", montant)
	time.Sleep(500 * time.Millisecond)

	if c.solde >= montant {
		c.solde -= montant
		fmt.Printf("Retrait réussi. Nouveau solde: %d$\n", c.solde)
		return true
	}

	fmt.Printf("❌ Solde insuffisant! Solde actuel: %d$\n", c.solde)
	return false
}

func (c *CompteBancaire) ObtenirSolde() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.solde
}

func main() {
	compte := CompteBancaire{solde: 1000}
	var wg sync.WaitGroup

	fmt.Println("========== COMPTE BANCAIRE ==========")
	fmt.Printf("Solde initial: %d$\n\n", compte.ObtenirSolde())

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			compte.Deposer(200 * id)
		}(i)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			compte.Retirer(150 * id)
		}(i)
	}

	wg.Wait()

	fmt.Println("\n========== RÉSUMÉ FINAL ==========")
	fmt.Printf("Solde final: %d$\n", compte.ObtenirSolde())
}
