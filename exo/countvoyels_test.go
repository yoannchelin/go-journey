package exo

import "testing"

func TestVoyels(t *testing.T) {
	tests := []struct {
		name string
		in string
		want int
	} {
		{"normal", "aeiouy", 6},
		{"number of voyels is %d", "", 0},
		{"number of voyels is %d", "la zone", 3},
		{"number of voyels is %d", "222222222", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Voyels(tc.in)
			if got != tc.want {
				t.Errorf("Voyels(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}

}