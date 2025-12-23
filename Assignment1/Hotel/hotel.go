package Hotel

import "fmt"

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Name  string
	Rooms map[string]Room
}

func NewHotel(name string) *Hotel {
	return &Hotel{
		Name:  name,
		Rooms: make(map[string]Room),
	}
}

func (h *Hotel) AddRoom(roomNumber, roomType string, price float64) error {
	if _, exists := h.Rooms[roomNumber]; exists {
		return fmt.Errorf("room %s already exists", roomNumber)
	}

	h.Rooms[roomNumber] = Room{
		RoomNumber:    roomNumber,
		Type:          roomType,
		PricePerNight: price,
		IsOccupied:    false,
	}
	return nil
}

func (h *Hotel) CheckIn(roomNumber string) error {
	room, exists := h.Rooms[roomNumber]
	if !exists {
		return fmt.Errorf("room %s does not exist", roomNumber)
	}

	if room.IsOccupied {
		return fmt.Errorf("room %s is already occupied", roomNumber)
	}

	room.IsOccupied = true
	h.Rooms[roomNumber] = room
	return nil
}

func (h *Hotel) CheckOut(roomNumber string) error {
	room, exists := h.Rooms[roomNumber]
	if !exists {
		return fmt.Errorf("room %s does not exist", roomNumber)
	}

	if !room.IsOccupied {
		return fmt.Errorf("room %s is already vacant", roomNumber)
	}

	room.IsOccupied = false
	h.Rooms[roomNumber] = room
	return nil
}

func (h *Hotel) ListVacantRooms() []Room {
	var vacantRooms []Room
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			vacantRooms = append(vacantRooms, room)
		}
	}
	return vacantRooms
}
