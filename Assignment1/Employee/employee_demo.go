package Employee

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDemo() {
	reader := bufio.NewReader(os.Stdin)
	var employees []Employee

	for {
		fmt.Println("\n1. Add Employee")
		fmt.Println("2. Show Salaries")
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
			fmt.Print("Type (1=FullTime, 2=PartTime, 3=Contractor, 4=Intern): ")
			typeStr, _ := reader.ReadString('\n')
			typeStr = strings.TrimSpace(typeStr)
			empType, _ := strconv.Atoi(typeStr)

			switch empType {
			case 1:
				fmt.Print("Monthly salary: ")
				salaryStr, _ := reader.ReadString('\n')
				salary, _ := strconv.ParseFloat(strings.TrimSpace(salaryStr), 64)

				fmt.Print("Bonus rate (0.15 for 15%): ")
				bonusStr, _ := reader.ReadString('\n')
				bonusRate, _ := strconv.ParseFloat(strings.TrimSpace(bonusStr), 64)

				employees = append(employees, FullTime{
					MonthlySalary: salary,
					BonusRate:     bonusRate,
				})
				fmt.Println("FullTime added")

			case 2:
				fmt.Print("Hourly rate: ")
				rateStr, _ := reader.ReadString('\n')
				hourlyRate, _ := strconv.ParseFloat(strings.TrimSpace(rateStr), 64)

				fmt.Print("Hours worked: ")
				hoursStr, _ := reader.ReadString('\n')
				hours, _ := strconv.Atoi(strings.TrimSpace(hoursStr))

				employees = append(employees, PartTime{
					HourlyRate:  hourlyRate,
					HoursWorked: hours,
				})
				fmt.Println("PartTime added")

			case 3:
				fmt.Print("Project rate: ")
				rateStr, _ := reader.ReadString('\n')
				projectRate, _ := strconv.ParseFloat(strings.TrimSpace(rateStr), 64)

				fmt.Print("Projects completed: ")
				projectsStr, _ := reader.ReadString('\n')
				projects, _ := strconv.Atoi(strings.TrimSpace(projectsStr))

				employees = append(employees, Contractor{
					ProjectRate:       projectRate,
					ProjectsCompleted: projects,
				})
				fmt.Println("Contractor added")

			case 4:
				fmt.Print("Daily rate: ")
				rateStr, _ := reader.ReadString('\n')
				dailyRate, _ := strconv.ParseFloat(strings.TrimSpace(rateStr), 64)

				fmt.Print("Days worked: ")
				daysStr, _ := reader.ReadString('\n')
				days, _ := strconv.Atoi(strings.TrimSpace(daysStr))

				employees = append(employees, Intern{
					DailyRate:  dailyRate,
					DaysWorked: days,
				})
				fmt.Println("Intern added")

			default:
				fmt.Println("Enter 1-4")
			}

		case 2:
			if len(employees) == 0 {
				fmt.Println("No employees")
				continue
			}

			total := 0.0
			for i, emp := range employees {
				salary := emp.CalculateSalary()
				bonus := emp.CalculateBonus()
				total += salary + bonus

				fmt.Printf("\nEmployee %d:\n", i+1)
				fmt.Printf("Salary: $%.2f\n", salary)
				fmt.Printf("Bonus: $%.2f\n", bonus)
				fmt.Printf("Total: $%.2f\n", salary+bonus)
			}
			fmt.Printf("\nAll employees total: $%.2f\n", total)

		case 3:
			return

		default:
			fmt.Println("Enter 1, 2, or 3")
		}
	}
}
