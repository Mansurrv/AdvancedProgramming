package Wallet

import (
	"fmt"
	"time"
)

type Transaction struct {
	Type        string
	Amount      float64
	Timestamp   time.Time
	Description string
}

type Wallet struct {
	Balance      float64
	Transactions []Transaction
	Owner        string
}

func NewWallet(owner string) *Wallet {
	return &Wallet{
		Balance:      0,
		Transactions: make([]Transaction, 0),
		Owner:        owner,
	}
}

func (w *Wallet) AddMoney(amount float64, description string) {
	if amount <= 0 {
		fmt.Println("Amount must be positive")
		return
	}

	w.Balance += amount
	transaction := Transaction{
		Type:        "CREDIT",
		Amount:      amount,
		Timestamp:   time.Now(),
		Description: description,
	}
	w.Transactions = append(w.Transactions, transaction)

	fmt.Printf("Added $%.2f to wallet. New balance: $%.2f\n", amount, w.Balance)
}

func (w *Wallet) SpendMoney(amount float64, description string) bool {
	if amount <= 0 {
		fmt.Println("Amount must be positive")
		return false
	}

	if amount > w.Balance {
		fmt.Printf("Insufficient funds. Balance: $%.2f, Required: $%.2f\n", w.Balance, amount)
		return false
	}

	w.Balance -= amount
	transaction := Transaction{
		Type:        "DEBIT",
		Amount:      amount,
		Timestamp:   time.Now(),
		Description: description,
	}
	w.Transactions = append(w.Transactions, transaction)

	fmt.Printf("Spent $%.2f. Remaining balance: $%.2f\n", amount, w.Balance)
	return true
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w *Wallet) ShowTransactions() {
	fmt.Printf("\n--- Transaction History for %s ---\n", w.Owner)
	if len(w.Transactions) == 0 {
		fmt.Println("No transactions yet")
		return
	}

	fmt.Printf("Current Balance: $%.2f\n\n", w.Balance)
	fmt.Println("Type\tAmount\t\tDate\t\t\tDescription")
	fmt.Println("----\t------\t\t----\t\t\t-----------")

	for _, t := range w.Transactions {
		fmt.Printf("%s\t$%.2f\t%s\t%s\n",
			t.Type,
			t.Amount,
			t.Timestamp.Format("2006-01-02 15:04"),
			t.Description)
	}
}
