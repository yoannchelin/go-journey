package warmup2

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		in int
		want bool
	} {
		{"le chiffres est pair", -4, true},
		{"le chiffres est impair", 3, false},
		{"le chiffres est pair", 8, true},
		{"le chiffres est pair", 0, true},
		{"le chiffres est impair", 1, false},

	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEven(tc.in)
			if got != tc.want {
				t.Errorf("IsEven(%d) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}

}
