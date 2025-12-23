package Gym

import (
	"fmt"
	"time"
)

type BasicMember struct {
	Name     string
	JoinDate time.Time
	IsActive bool
}

func (bm BasicMember) GetDetails() string {
	status := "Active"
	if !bm.IsActive {
		status = "Inactive"
	}
	return fmt.Sprintf("Type: Basic Member\nName: %s\nJoin Date: %s\nStatus: %s",
		bm.Name, bm.JoinDate.Format("2006-01-02"), status)
}

type PremiumMember struct {
	Name            string
	JoinDate        time.Time
	IsActive        bool
	PersonalTrainer string
	LockerNumber    int
}

func (pm PremiumMember) GetDetails() string {
	status := "Active"
	if !pm.IsActive {
		status = "Inactive"
	}
	return fmt.Sprintf("Type: Premium Member\nName: %s\nJoin Date: %s\nStatus: %s\nPersonal Trainer: %s\nLocker Number: %d",
		pm.Name, pm.JoinDate.Format("2006-01-02"), status, pm.PersonalTrainer, pm.LockerNumber)
}
