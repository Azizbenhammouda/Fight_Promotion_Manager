package fighters

import (
	"testing"
	"time"
)

func TestAgeOn(t *testing.T) {
	tests := []struct {
		name string
		dob  time.Time
		on   time.Time
		want int
	}{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AgeOn(tc.dob, tc.on)
			if got != tc.want {
				t.Errorf("AgeOn(%v, %v) = %d, want %d", tc.dob, tc.on, got, tc.want)
			}
		})
	}
}
