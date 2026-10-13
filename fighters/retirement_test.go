package fighters

import "testing"

func TestShouldRetire(t *testing.T) {
	tests := []struct {
		name string
		age  int
		roll float64
		want bool
	}{
		{"young never retires", 25, 0.0, false},
		{"early bracket, roll under chance", 34, 0.004, true},
		{"early bracket, roll at chance", 34, 0.005, false},
		{"mid bracket, roll under chance", 38, 0.019, true},
		{"old fighter, high roll", 45, 0.5, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ShouldRetire(tc.age, tc.roll)
			if got != tc.want {
				t.Errorf("ShouldRetire(%d, %v) = %v, want %v", tc.age, tc.roll, got, tc.want)
			}
		})
	}
}
