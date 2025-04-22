package main

import (
	"fmt"

	"github.com/cetinboran/earn_sandwich/models"
)

func main() {
	pool := models.NewPool()

	// Base price of asset before attack
	basePrice := 100.0

	// 1. Find an in-progress victim transaction
	inProgressTxs := pool.GetTransactionsByStatus("in-progress")
	if len(inProgressTxs) == 0 {
		fmt.Println("No in-progress transactions found.")
		return
	}

	victim := inProgressTxs[0]
	fmt.Println("🎯 Victim Transaction Found:")
	fmt.Printf("ID: %s, From: %s, To: %s, Amount: %.2f\n", victim.ID, victim.From, victim.To, victim.Amount)

	// 2. Execute Sandwich Attack
	pool.SandwichAttack(victim, basePrice)

	// 3. Show pool after attack
	fmt.Println("\n📦 Transaction Pool After Attack:")
	for _, tx := range pool.Transactions {
		fmt.Printf("ID: %s, From: %s, To: %s, Amount: %.2f, Type: %s, Status: %s\n",
			tx.ID, tx.From, tx.To, tx.Amount, tx.Type, tx.Status)
	}
}
