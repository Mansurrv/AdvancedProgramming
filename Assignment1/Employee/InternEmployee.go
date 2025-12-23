package employee

type InternEmployee struct {
	DailyRate  float64
	DaysWorked int
}

func (i InternEmployee) CalculateSalary() float64 {
	return i.DailyRate * float64(i.DaysWorked)
}

func (i InternEmployee) CalculateBonus() float64 {
	return float64(i.DaysWorked) * 10
}

func (i InternEmployee) GetWorkHours() int {
	return i.DaysWorked * 4
}
