package Gym

import (
	"fmt"
)

type Member interface {
	GetDetails() string
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() *Gym {
	return &Gym{
		Members: make(map[uint64]Member),
	}
}

func (g *Gym) AddMember(id uint64, member Member) {
	if _, exists := g.Members[id]; exists {
		fmt.Printf("Member with ID %d already exists\n", id)
		return
	}
	g.Members[id] = member
	fmt.Printf("Member with ID %d added successfully\n", id)
}

func (g *Gym) ListMembers() {
	fmt.Println("\n--- Gym Membership List ---")
	if len(g.Members) == 0 {
		fmt.Println("No members registered")
		return
	}
	for id, member := range g.Members {
		fmt.Printf("ID: %d\n%s\n", id, member.GetDetails())
	}
}
