package Hotel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDemo() {
	reader := bufio.NewReader(os.Stdin)
	hotel := NewHotel("Demo Hotel")

	for {
		fmt.Println("\n1. Add Room")
		fmt.Println("2. Check In")
		fmt.Println("3. Check Out")
		fmt.Println("4. Show Vacant Rooms")
		fmt.Println("5. Back to Main")
		fmt.Print("Choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Enter 1-5")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Room number: ")
			roomNum, _ := reader.ReadString('\n')
			roomNum = strings.TrimSpace(roomNum)

			fmt.Print("Room type: ")
			roomType, _ := reader.ReadString('\n')
			roomType = strings.TrimSpace(roomType)

			fmt.Print("Price per night: ")
			priceStr, _ := reader.ReadString('\n')
			priceStr = strings.TrimSpace(priceStr)
			price, _ := strconv.ParseFloat(priceStr, 64)

			err := hotel.AddRoom(roomNum, roomType, price)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Room added")
			}

		case 2:
			fmt.Print("Room number to check in: ")
			roomNum, _ := reader.ReadString('\n')
			roomNum = strings.TrimSpace(roomNum)

			err := hotel.CheckIn(roomNum)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Checked in")
			}

		case 3:
			fmt.Print("Room number to check out: ")
			roomNum, _ := reader.ReadString('\n')
			roomNum = strings.TrimSpace(roomNum)

			err := hotel.CheckOut(roomNum)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Checked out")
			}

		case 4:
			rooms := hotel.ListVacantRooms()
			if len(rooms) == 0 {
				fmt.Println("No vacant rooms")
			} else {
				fmt.Printf("\nVacant rooms (%d):\n", len(rooms))
				for _, room := range rooms {
					fmt.Printf("Room %s: %s - $%.2f\n",
						room.RoomNumber, room.Type, room.PricePerNight)
				}
			}

		case 5:
			return

		default:
			fmt.Println("Enter 1-5")
		}
	}
}
