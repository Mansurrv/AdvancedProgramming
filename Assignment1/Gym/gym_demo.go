package Gym

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDemo() {
	reader := bufio.NewReader(os.Stdin)
	gym := NewGym("My Gym")

	for {
		fmt.Println("\n1. Add Member")
		fmt.Println("2. List Members")
		fmt.Println("3. Back to Main")
		fmt.Print("Choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Enter 1, 2, or 3")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Member name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Type (1=Basic, 2=Premium): ")
			typeStr, _ := reader.ReadString('\n')
			typeStr = strings.TrimSpace(typeStr)
			memberType, _ := strconv.Atoi(typeStr)

			fmt.Print("Duration (months): ")
			durationStr, _ := reader.ReadString('\n')
			durationStr = strings.TrimSpace(durationStr)
			duration, _ := strconv.Atoi(durationStr)

			if memberType == 1 {
				gym.AddMember(BasicMember{
					Name:               name,
					MembershipDuration: duration,
				})
				fmt.Println("Basic member added")
			} else if memberType == 2 {
				fmt.Print("Personal trainer? (y/n): ")
				trainerStr, _ := reader.ReadString('\n')
				trainerStr = strings.TrimSpace(strings.ToLower(trainerStr))
				hasTrainer := trainerStr == "y"

				gym.AddMember(PremiumMember{
					Name:               name,
					MembershipDuration: duration,
					PersonalTrainer:    hasTrainer,
				})
				fmt.Println("Premium member added")
			} else {
				fmt.Println("Enter 1 or 2")
			}

		case 2:
			members := gym.GetMembers()
			if len(members) == 0 {
				fmt.Println("No members")
			} else {
				basic, premium := gym.GetMemberCounts()
				fmt.Printf("\nMembers: %d total (%d basic, %d premium)\n",
					len(members), basic, premium)

				for id, member := range members {
					fmt.Printf("ID %d: %s\n", id, member.GetDetails())
				}
			}

		case 3:
			return

		default:
			fmt.Println("Enter 1, 2, or 3")
		}
	}
}
