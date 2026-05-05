package math

import (
	"errors"
	"testing"
)

func TestPlus(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"deux positifs", 2, 3, 5},
		{"avec zéro", 0, 7, 7},
		{"négatifs", -2, -3, -5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Plus(tc.a, tc.b); got != tc.want {
				t.Errorf("Plus(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestPlusPlus(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		want    int
	}{
		{"positifs", 1, 2, 3, 6},
		{"avec zéros", 0, 0, 5, 5},
		{"négatifs", -1, -2, -3, -6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := PlusPlus(tc.a, tc.b, tc.c); got != tc.want {
				t.Errorf("PlusPlus(%d,%d,%d) = %d, want %d", tc.a, tc.b, tc.c, got, tc.want)
			}
		})
	}
}

func TestSumAny(t *testing.T) {
	t.Run("cas heureux", func(t *testing.T) {
		sum2, sum3, err := SumAny(1, 2, 3)
		if err != nil {
			t.Fatalf("attendait pas d'erreur, got %v", err)
		}
		if sum2 != 3 {
			t.Errorf("sum2 = %d, want 3", sum2)
		}
		if sum3 != 6 {
			t.Errorf("sum3 = %d, want 6", sum3)
		}
	})

	t.Run("a invalide", func(t *testing.T) {
		_, _, err := SumAny("foo", 2, 3)
		if !errors.Is(err, ErrInvalidNumber) {
			t.Errorf("attendait ErrInvalidNumber, got %v", err)
		}
	})

	t.Run("c est nil", func(t *testing.T) {
		_, _, err := SumAny(1, 2, nil)
		if !errors.Is(err, ErrInvalidNumber) {
			t.Errorf("attendait ErrInvalidNumber, got %v", err)
		}
	})
}
