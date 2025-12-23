package employee

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}

func (p PartTime) CalculateBonus() float64 {
	if p.HoursWorked > 100 {
		return p.CalculateSalary() * 0.05
	}
	return 0
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}
