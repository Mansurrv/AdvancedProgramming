package hotel

import "fmt"

type Room struct {
	RoomNumber    string
	RoomType      string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[string]Room
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
	fmt.Println("Room added successfully")
}

func (h *Hotel) CheckIn(roomNumber string) {
	room, exists := h.Rooms[roomNumber]
	if !exists {
		fmt.Println("")
		fmt.Println("Room doesn't exist!")
		return
	}
	if room.IsOccupied {
		fmt.Println("")
		fmt.Println("Room is already occupied")
		return
	}
}

func (h *Hotel) CheckOut(roomNumber string) {
	room, exists := h.Rooms[roomNumber]
	if !exists {
		fmt.Println("")
		fmt.Println("Room doesn't exist")
		return
	}
	if !room.IsOccupied {
		fmt.Println("")
		fmt.Println("Room is already vacant")
		return
	}
	room.IsOccupied = false
	h.Rooms[roomNumber] = room
	fmt.Println("Check out successful")
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant Rooms:")
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Printf("Room: %s | Type: %s | Price: $%.2f\n", room.RoomNumber, room.RoomType, room.PricePerNight)
		}
	}
}
