package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulate price impact based on transaction type
func simulateMarketImpact(basePrice float64, tx Transaction) float64 {
	switch tx.Type {
	case "front-run":
		return basePrice + tx.Amount*0.04 // bot buys early -> price increases more
	case "normal":
		return basePrice + tx.Amount*0.02 // user swaps -> price increases moderately
	case "back-run":
		return basePrice - tx.Amount*0.03 // bot sells after -> price decreases
	default:
		return basePrice
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	basePrice := 100.0
	fmt.Println("📊 Initial token price:", basePrice)

	// Define transactions
	frontRunTx := Transaction{
		From:   "Bot",
		To:     "DEX",
		Amount: 15,
		Type:   "front-run",
	}

	userTx := Transaction{
		From:   "User",
		To:     "DEX",
		Amount: 10,
		Type:   "normal",
	}

	backRunTx := Transaction{
		From:   "DEX",
		To:     "Bot",
		Amount: 15,
		Type:   "back-run",
	}

	fmt.Println("\n🚨 Sandwich Attack Simulation Starting...")

	// Step 1: Bot performs front-run (buys before user)
	fmt.Println("\n1️⃣ Front-run Bot Buys Before User")
	fmt.Printf("   Bot buys at price: $%.2f\n", basePrice)
	basePrice = simulateMarketImpact(basePrice, frontRunTx)

	// Step 2: User performs normal transaction at increased price
	fmt.Println("2️⃣ User Swaps Tokens (Pays Higher Price)")
	fmt.Printf("   User buys at price: $%.2f\n", basePrice)
	basePrice = simulateMarketImpact(basePrice, userTx)

	// Step 3: Bot sells after user (back-run)
	fmt.Println("3️⃣ Bot Sells After User")
	fmt.Printf("   Bot sells at price: $%.2f\n", basePrice)
	sellPrice := basePrice
	basePrice = simulateMarketImpact(basePrice, backRunTx)

	// Calculate bot's profit
	buyPrice := 100.0 + frontRunTx.Amount*0.04 // original price + front-run impact
	profit := (sellPrice - buyPrice) * frontRunTx.Amount
	fmt.Printf("   💰 Bot profit from sandwich: $%.2f\n", profit)
}
