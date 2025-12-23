package Wallet

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDemo() {
	reader := bufio.NewReader(os.Stdin)
	wallet := NewWallet()

	for {
		fmt.Println("\n1. Add Money")
		fmt.Println("2. Spend Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Back to Main")
		fmt.Print("Choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Enter 1-4")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Amount to add: $")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, _ := strconv.ParseFloat(amountStr, 64)

			fmt.Print("Note: ")
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			wallet.AddMoney(amount, note)
			fmt.Printf("Added $%.2f\n", amount)

		case 2:
			fmt.Print("Amount to spend: $")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, _ := strconv.ParseFloat(amountStr, 64)

			fmt.Print("Note: ")
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			success := wallet.SpendMoney(amount, note)
			if success {
				fmt.Printf("Spent $%.2f\n", amount)
			} else {
				fmt.Printf("Cannot spend $%.2f (balance: $%.2f)\n",
					amount, wallet.GetBalance())
			}

		case 3:
			balance := wallet.GetBalance()
			transactions := wallet.GetTransactions()

			fmt.Printf("\nBalance: $%.2f\n", balance)
			fmt.Printf("Transactions: %d\n", len(transactions))

			if len(transactions) > 0 {
				fmt.Println("Recent:")
				start := 0
				if len(transactions) > 3 {
					start = len(transactions) - 3
				}
				for i := start; i < len(transactions); i++ {
					t := transactions[i]
					fmt.Printf("  %s: $%.2f - %s\n", t.Type, t.Amount, t.Note)
				}
			}

		case 4:
			return

		default:
			fmt.Println("Enter 1-4")
		}
	}
}
