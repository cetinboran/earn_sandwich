# 🥪 Sandwich Attack Simulation — Go-based Front-Running Scenario in DeFi

## 🎯 Lab Objective

This simulation demonstrates how a **sandwich attack** works in a decentralized finance (DeFi) environment, particularly focusing on **front-running and back-running of transactions**. Though not executed on a real blockchain, this lab accurately models the logic behind sandwich attacks to help developers, security researchers, and blockchain enthusiasts better understand these threats.

> ✅ This is a fully **offline simulation** — no real blockchain, no wallets, no tokens. It's purely educational, providing a safe environment to grasp the core concept of how sandwich attacks operate.

---

## ⚒️ Technical Summary

- **Attack Type**: Sandwich Attack (Front-running + Back-running)
- **Environment**: Go (Golang)
- **Goal**: Simulate a victim’s transaction being sandwiched between two attacker transactions
- **Realism**: Fully simulated (no real token transfers or blockchain interaction)
- **Educational Focus**: Observe attacker profit through price manipulation during a simulated swap

---

## 📦 Project Structure

| File        | Description                                                         |
| ----------- | ------------------------------------------------------------------- |
| `main.go`   | Core simulation logic for modeling victim and attacker transactions |
| `README.md` | Project documentation (you’re reading it!)                          |

---

## 🚀 Getting Started

### ✅ Prerequisites

- [Go](https://go.dev/doc/install) (version 1.19 or later)

### 🛠️ Running the Simulation

Clone the repository and run the simulation:

```bash
go run main.go
```

The output will describe the sequence of events:

1. The pool starts with a fixed amount of token A and B.
2. The victim tries to swap a token.
3. The attacker front-runs with a transaction that moves the price.
4. The victim's transaction executes at a worse price.
5. The attacker back-runs to reverse their position and profit.

---

## 🧪 How the Sandwich Attack Works

### 🍞 Normal Swap (No Attack)

```go
// Victim wants to swap 10 tokens of A for B
victimTx := Transaction{From: "victim", AmountIn: 10, Direction: "AtoB"}
```

Without interference, the victim gets the best market rate.

---

### 🥪 Simulated Sandwich Flow

```go
// Attacker steps in
frontRunTx := Transaction{From: "attacker", AmountIn: 30, Direction: "AtoB"}

// Victim follows (sandwiched)
victimTx := Transaction{From: "victim", AmountIn: 10, Direction: "AtoB"}

// Attacker exits (back-runs)
backRunTx := Transaction{From: "attacker", AmountIn: 30, Direction: "BtoA"}
```

The attacker manipulates the pool rate with a large transaction, causing the victim to swap at a worse rate. Then, the attacker profits by executing the reverse trade.

---

## 📊 Sample Output

```text
Initial Pool State: A: 1000.00 | B: 1000.00
[FRONT-RUN] Attacker swapped 30.00 A → 22.56 B
[VICTIM] Victim swapped 10.00 A → 6.15 B
[BACK-RUN] Attacker swapped 30.00 B → 38.59 A
[PROFIT] Attacker earned: +8.59 A (net)
```

You can clearly observe the attacker’s profit gained by exploiting the price slippage caused by their own front-running move.

---

## 🔒 What This Teaches

This simulation shows how:

- **Transaction ordering** can be abused in mempool-based blockchain systems
- **Slippage tolerance** impacts DeFi users
- **Front-running** and **back-running** are practical and profitable
- Such attacks are often executed by **bots** or MEV strategies in real-world scenarios

---

## 🛡️ How to Prevent Sandwich Attacks

1. **Use Slippage Limits**  
   Set strict slippage tolerance in your swap transactions.

2. **Enable Transaction Privacy**  
   Use tools like Flashbots or private relays that prevent public mempool exposure.

3. **Batch Transactions**  
   Bundle transactions to avoid being front-ran between submission and confirmation.

4. **MEV-aware Design**  
   Use fair sequencing or encrypted mempools if available in your blockchain infrastructure.

---

## 🧠 Final Thoughts

By simulating a sandwich attack in Go, we provide an approachable way to understand a real and widespread DeFi exploit without deploying contracts or interacting with blockchains. This approach empowers both developers and security researchers to test, observe, and reason about these attacks safely.

---

## 📚 References

- [Golang Docs](https://go.dev/doc/)
