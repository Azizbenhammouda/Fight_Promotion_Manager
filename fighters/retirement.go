package fighters

const (
	earlyRetirementAge = 33
	midRetirementAge   = 36
	lateRetirementAge  = 40

	chanceEarly = 0.005
	chanceMid   = 0.02
	chanceLate  = 0.08
)

func RetirementChance(age int) float64 {
	switch {
	case age < earlyRetirementAge:
		return 0
	case age < midRetirementAge:
		return chanceEarly
	case age < lateRetirementAge:
		return chanceMid
	default:
		return chanceLate
	}
}

func ShouldRetire(age int, roll float64) bool {
	return roll < RetirementChance(age)
}
