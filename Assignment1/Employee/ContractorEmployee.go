package employee

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return c.ProjectRate * float64(c.ProjectsCompleted)
}

func (c Contractor) CalculateBonus() float64 {
	if c.ProjectsCompleted > 5 {
		return float64(c.ProjectsCompleted) * 100
	}
	return 0
}

func (c Contractor) GetWorkHours() int {
	return c.ProjectsCompleted * 20
}
