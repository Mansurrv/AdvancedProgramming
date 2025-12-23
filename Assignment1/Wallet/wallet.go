package Wallet

import "time"

type TransactionType string

const (
	Deposit    TransactionType = "DEPOSIT"
	Withdrawal TransactionType = "WITHDRAWAL"
)

type Transaction struct {
	ID        int
	Type      TransactionType
	Amount    float64
	Timestamp time.Time
	Note      string
}

type Wallet struct {
	Balance      float64
	Transactions []Transaction
	nextID       int
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance:      0,
		Transactions: make([]Transaction, 0),
		nextID:       1,
	}
}

func (w *Wallet) AddMoney(amount float64, note string) {
	w.Balance += amount
	transaction := Transaction{
		ID:        w.nextID,
		Type:      Deposit,
		Amount:    amount,
		Timestamp: time.Now(),
		Note:      note,
	}
	w.Transactions = append(w.Transactions, transaction)
	w.nextID++
}

func (w *Wallet) SpendMoney(amount float64, note string) bool {
	if amount > w.Balance {
		return false
	}

	w.Balance -= amount
	transaction := Transaction{
		ID:        w.nextID,
		Type:      Withdrawal,
		Amount:    amount,
		Timestamp: time.Now(),
		Note:      note,
	}
	w.Transactions = append(w.Transactions, transaction)
	w.nextID++
	return true
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w *Wallet) GetTransactions() []Transaction {
	return w.Transactions
}
