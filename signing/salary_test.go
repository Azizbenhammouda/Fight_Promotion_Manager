package signing

import (
	"testing"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
)

func TestMonthlySalary(t *testing.T) {
	tests := []struct {
		name       string
		skills     fighters.FightingSkills
		popularity int
		reputation int
		want       int64
	}{
		{"start of game", fighters.FightingSkills{Wrestling: 59}, 0, 5, 2212},
		{"big promotion", fighters.FightingSkills{Wrestling: 59}, 0, 10, 2950},
		{"star fighter", fighters.FightingSkills{Wrestling: 59}, 100, 10, 5900},
		{"factor is capped", fighters.FightingSkills{Wrestling: 59}, 0, 100, 4425},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MonthlySalary(tc.skills, tc.popularity, tc.reputation)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
