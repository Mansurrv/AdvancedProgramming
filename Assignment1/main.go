package main

import (
	employee "Assignment1/Employee"
	gym "Assignment1/Gym"
	hotel "Assignment1/Hotel"
	wallet "Assignment1/Wallet"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	h := hotel.Hotel{Rooms: make(map[string]hotel.Room)}
	g := gym.NewGym()
	w := wallet.NewWallet("John Doe")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Management System ===")
		fmt.Println("1. Hotel Management")
		fmt.Println("2. Gym Management")
		fmt.Println("3. Wallet Management")
		fmt.Println("4. Employee Details")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			hotelManagement(&h, scanner)
		case "2":
			gymManagement(g, scanner)
		case "3":
			walletManagement(w, scanner)
		case "4":
			showEmployeeDetails()
		case "5":
			fmt.Println("\nExiting...")
			return
		default:
			fmt.Println("\nInvalid choice. Try again.")
		}
	}
}

func hotelManagement(h *hotel.Hotel, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Hotel Management ---")
		fmt.Println("1. Add Room")
		fmt.Println("2. Check In")
		fmt.Println("3. Check Out")
		fmt.Println("4. List Vacant Rooms")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			var roomNumber, roomType string
			var price float64

			fmt.Print("Enter room number: ")
			scanner.Scan()
			roomNumber = scanner.Text()

			fmt.Print("Enter room type: ")
			scanner.Scan()
			roomType = scanner.Text()

			fmt.Print("Enter price per night: ")
			scanner.Scan()
			priceStr := scanner.Text()
			price, _ = strconv.ParseFloat(priceStr, 64)

			r := hotel.Room{
				RoomNumber:    roomNumber,
				RoomType:      roomType,
				PricePerNight: price,
				IsOccupied:    false,
			}

			h.AddRoom(r)

		case "2":
			fmt.Print("Enter room number to check in: ")
			scanner.Scan()
			roomNumber := scanner.Text()
			h.CheckIn(roomNumber)

		case "3":
			fmt.Print("Enter room number to check out: ")
			scanner.Scan()
			roomNumber := scanner.Text()
			h.CheckOut(roomNumber)

		case "4":
			h.ListVacantRooms()

		case "5":
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func gymManagement(g *gym.Gym, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Gym Management ---")
		fmt.Println("1. Add Basic Member")
		fmt.Println("2. Add Premium Member")
		fmt.Println("3. List All Members")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			fmt.Print("Enter member name: ")
			scanner.Scan()
			name := scanner.Text()

			basicMember := gym.BasicMember{
				Name:     name,
				JoinDate: time.Now(),
				IsActive: true,
			}

			g.AddMember(id, basicMember)

		case "2":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			fmt.Print("Enter member name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter personal trainer name: ")
			scanner.Scan()
			trainer := scanner.Text()

			fmt.Print("Enter locker number: ")
			scanner.Scan()
			lockerStr := scanner.Text()
			lockerNum, err := strconv.Atoi(lockerStr)
			if err != nil {
				fmt.Println("Invalid locker number. Please enter a number.")
				continue
			}

			premiumMember := gym.PremiumMember{
				Name:            name,
				JoinDate:        time.Now(),
				IsActive:        true,
				PersonalTrainer: trainer,
				LockerNumber:    lockerNum,
			}

			g.AddMember(id, premiumMember)

		case "3":
			g.ListMembers()

		case "4":
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func walletManagement(w *wallet.Wallet, scanner *bufio.Scanner) {
	for {
		fmt.Printf("\n--- Wallet Management (Owner: %s) ---\n", w.Owner)
		fmt.Printf("Current Balance: $%.2f\n", w.GetBalance())
		fmt.Println("1. Add Money")
		fmt.Println("2. Spend Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. View Transactions")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter amount to add: $")
			scanner.Scan()
			amountStr := scanner.Text()
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}

			fmt.Print("Enter description: ")
			scanner.Scan()
			description := scanner.Text()

			w.AddMoney(amount, description)

		case "2":
			fmt.Print("Enter amount to spend: $")
			scanner.Scan()
			amountStr := scanner.Text()
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}

			fmt.Print("Enter description: ")
			scanner.Scan()
			description := scanner.Text()

			w.SpendMoney(amount, description)

		case "3":
			fmt.Printf("\nCurrent Balance: $%.2f\n", w.GetBalance())

		case "4":
			w.ShowTransactions()

		case "5":
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func showEmployeeDetails() {
	fmt.Println("\n--- Employee Details ---")
	employees := []employee.Employee{
		employee.FullTime{
			MonthlySalary: 5000,
			BonusRate:     0.1,
			WorkHours:     100,
		},
		employee.PartTime{
			HourlyRate:  20,
			HoursWorked: 120,
		},
		employee.Contractor{
			ProjectRate:       800,
			ProjectsCompleted: 7,
		},
		employee.InternEmployee{
			DailyRate:  50,
			DaysWorked: 30,
		},
	}

	for _, e := range employees {
		fmt.Println("Employee Details: ")
		employee.PrintEmployeeDetails(e)
		fmt.Println()
	}
}
