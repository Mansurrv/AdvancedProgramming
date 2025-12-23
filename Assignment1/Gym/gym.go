package Gym

type Member interface {
	GetDetails() string
	GetType() string
	GetID() uint64
}

type BasicMember struct {
	ID                 uint64
	Name               string
	MembershipDuration int
}

func (b BasicMember) GetDetails() string {
	return "Basic Member"
}

func (b BasicMember) GetType() string {
	return "basic"
}

func (b BasicMember) GetID() uint64 {
	return b.ID
}

type PremiumMember struct {
	ID                 uint64
	Name               string
	MembershipDuration int
	PersonalTrainer    bool
}

func (p PremiumMember) GetDetails() string {
	return "Premium Member"
}

func (p PremiumMember) GetType() string {
	return "premium"
}

func (p PremiumMember) GetID() uint64 {
	return p.ID
}

type Gym struct {
	Name    string
	Members map[uint64]Member
	nextID  uint64
}

func NewGym(name string) *Gym {
	return &Gym{
		Name:    name,
		Members: make(map[uint64]Member),
		nextID:  1,
	}
}

func (g *Gym) AddMember(member Member) {
	memberID := g.nextID

	switch m := member.(type) {
	case BasicMember:
		m.ID = memberID
		g.Members[memberID] = m
	case PremiumMember:
		m.ID = memberID
		g.Members[memberID] = m
	}

	g.nextID++
}

func (g *Gym) GetMembers() map[uint64]Member {
	return g.Members
}

func (g *Gym) GetMemberCounts() (int, int) {
	basicCount := 0
	premiumCount := 0

	for _, member := range g.Members {
		if member.GetType() == "basic" {
			basicCount++
		} else {
			premiumCount++
		}
	}

	return basicCount, premiumCount
}
