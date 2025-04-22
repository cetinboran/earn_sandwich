package models

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
)

type Transaction struct {
	ID     uuid.UUID
	From   string
	To     string
	Amount float64
	Type   string
	Status string
}

type Pool struct {
	Transactions []Transaction
}

func (p *Pool) AddTransaction(tx Transaction) {
	p.Transactions = append(p.Transactions, tx)
}

func (p *Pool) GetTransactionsByStatus(status string) []Transaction {
	var filteredTransactions []Transaction
	for _, tx := range p.Transactions {
		if tx.Status == status {
			filteredTransactions = append(filteredTransactions, tx)
		}
	}
	return filteredTransactions
}

func (p *Pool) GetTransactionWithID(id string) *Transaction {
	for i := range p.Transactions {
		if p.Transactions[i].ID.String() == id {
			return &p.Transactions[i]
		}
	}
	return nil
}

func (p *Pool) AddDummyData() {
	dummyTransactions := []Transaction{
		{ID: uuid.New(), From: "Bot1", To: "DEX1", Amount: rand.Float64() * 100, Type: "front-run", Status: "pending"},
		{ID: uuid.New(), From: "User1", To: "DEX1", Amount: rand.Float64() * 50, Type: "normal", Status: "in-progress"},
		{ID: uuid.New(), From: "Bot2", To: "DEX2", Amount: rand.Float64() * 80, Type: "back-run", Status: "completed"},
		{ID: uuid.New(), From: "User2", To: "DEX2", Amount: rand.Float64() * 200, Type: "normal", Status: "completed"},
		{ID: uuid.New(), From: "Bot3", To: "DEX3", Amount: rand.Float64() * 150, Type: "front-run", Status: "pending"},
	}

	for _, tx := range dummyTransactions {
		p.AddTransaction(tx)
	}
}

func NewPool() *Pool {
	pool := &Pool{Transactions: []Transaction{}}
	pool.AddDummyData()
	return pool
}

// SandwichAttack performs a sandwich attack on a victim transaction.
func (p *Pool) SandwichAttack(victim Transaction, basePrice float64) {
	// Front-run: buy at base price
	frontRunAmount := victim.Amount * 1.1
	frontRun := Transaction{
		ID:     uuid.New(),
		From:   "SandwichBot",
		To:     victim.To,
		Amount: frontRunAmount,
		Type:   "front-run",
		Status: "pending",
	}

	// Price increases after victim transaction (simulate slippage)
	newPrice := basePrice * 1.05

	// Back-run: sell at new price
	backRunAmount := frontRunAmount // Assume selling same quantity bought
	backRun := Transaction{
		ID:     uuid.New(),
		From:   "SandwichBot",
		To:     victim.To,
		Amount: backRunAmount,
		Type:   "back-run",
		Status: "pending",
	}

	// Calculate profit
	cost := frontRunAmount * basePrice
	revenue := backRunAmount * newPrice
	profit := revenue - cost

	fmt.Println("\n💥 Sandwich Attack Executed!")
	fmt.Printf("Victim ID: %s\n", victim.ID)
	fmt.Printf("Front-run Buy: %.2f at $%.2f = $%.2f\n", frontRunAmount, basePrice, cost)
	fmt.Printf("Back-run Sell: %.2f at $%.2f = $%.2f\n", backRunAmount, newPrice, revenue)
	fmt.Printf("💰 Profit: $%.2f\n", profit)

	p.AddTransaction(frontRun)
	p.AddTransaction(backRun)
}
