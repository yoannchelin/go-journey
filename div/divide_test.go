package div

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name          string
		a, b          int
		wantQuotient  int
		wantRemainder int
		wantErr       bool
	}{
		{"division normale", 10, 3, 3, 1, false},
		{"division exacte", 9, 3, 3, 0, false},
		{"a égal zéro", 0, 5, 0, 0, false},
		{"résultat négatif", -10, 3, -3, -1, false},
		{"diviseur négatif", 10, -3, -3, 1, false},
		{"division par zéro", 10, 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quotient, remainder, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Errorf("attendait une erreur, n'en a pas eu")
				}
				return
			}

			if err != nil {
				t.Errorf("erreur inattendue : %v", err)
			}
			if quotient != tt.wantQuotient {
				t.Errorf("quotient = %d, attendait %d", quotient, tt.wantQuotient)
			}
			if remainder != tt.wantRemainder {
				t.Errorf("reste = %d, attendait %d", remainder, tt.wantRemainder)
			}
		})
	}
}
