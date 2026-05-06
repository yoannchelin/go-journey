package exo

import "testing"

func TestCountVoyels(t *testing.T) {
	tests := []struct {
		name string
		in string
		want int
	} {
		{"number of voyels is %d", "aeiouy", 6},
		{"number of voyels is %d", "", 0},
		{"number of voyels is %d", "la zone", 3},
		{"number of voyels is %d", "222222222", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountVowels(tc.in)
			if got != tc.want {
				t.Errorf("CountVowels(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}