package employee

import "fmt"

type Employee interface {
	CalculateSalary() float64
	CalculateBonus() float64
	GetWorkHours() int
}

func PrintEmployeeDetails(e Employee) {
	fmt.Println("Salary: ", e.CalculateSalary())
	fmt.Println("Bonus: ", e.CalculateBonus())
	fmt.Println("Work Hours: ", e.GetWorkHours())
	fmt.Println("----------------------")
}
