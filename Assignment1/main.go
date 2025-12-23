package main

import (
	"fmt"
	"os"

	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
)

func main() {
	for {
		fmt.Println("\n=== Assignment 1 ===")
		fmt.Println("1. Hotel Room Reservation System")
		fmt.Println("2. Employee Salary Calculator")
		fmt.Println("3. Gym Membership Management")
		fmt.Println("4. Wallet Transaction Simulation")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Hotel.RunDemo()
		case 2:
			Employee.RunDemo()
		case 3:
			Gym.RunDemo()
		case 4:
			Wallet.RunDemo()
		case 5:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
