package employee

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
	WorkHours     int
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary + f.BonusRate
}

func (f FullTime) GetWorkHours() int {
	return f.WorkHours
}
