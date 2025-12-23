package Employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
	CalculateBonus() float64
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary
}

func (f FullTime) GetWorkHours() int {
	return 160
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary * f.BonusRate
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}

func (p PartTime) CalculateBonus() float64 {
	if p.HoursWorked > 80 {
		return p.CalculateSalary() * 0.10
	}
	return 0
}

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return c.ProjectRate * float64(c.ProjectsCompleted)
}

func (c Contractor) GetWorkHours() int {
	return c.ProjectsCompleted * 20
}

func (c Contractor) CalculateBonus() float64 {
	if c.ProjectsCompleted > 5 {
		return float64(c.ProjectsCompleted-5) * 100
	}
	return 0
}

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (i Intern) CalculateSalary() float64 {
	return i.DailyRate * float64(i.DaysWorked)
}

func (i Intern) GetWorkHours() int {
	return i.DaysWorked * 8
}

func (i Intern) CalculateBonus() float64 {
	if i.DaysWorked > 20 {
		return i.DailyRate * 2
	}
	return 0
}
