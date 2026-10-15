package signing

import "github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"

const (
	salaryPerSkillPoint = 50
	popularityDivisor   = 100.0
	promotionBaseFactor = 0.5
	reputationDivisor   = 20.0
	maxPromotionFactor  = 1.5
)

func bestSkill(s fighters.FightingSkills) int {
	return max(s.Boxing, s.KickBoxing, s.MuayThai, s.Wrestling, s.BJJ, s.Judo)
}
func MonthlySalary(skills fighters.FightingSkills, popularity, reputation int) int64 {
	base := float64(bestSkill(skills) * salaryPerSkillPoint)
	popularityFactor := 1 + float64(popularity)/popularityDivisor
	promotionFactor := min(promotionBaseFactor+float64(reputation)/reputationDivisor, maxPromotionFactor)
	return int64(base * popularityFactor * promotionFactor)
}
