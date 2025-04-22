package main

// Developed by Çetin Boran Mesüm

type Transaction struct {
	From   string
	To     string
	Amount float64
	Type   string // "normal", "front-run", "back-run"
}
