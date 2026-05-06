package exo

import (
	"testing" 
	"errors"
)

func TestMinSlice(t *testing.T)  {
	tests := []struct {
		name string
		in []int
		want int
		wantErr bool
	} {
		{"the smallest value is", []int{3, 1, 4, 1, 5, 9, 2, 6}, 1, false},
		{"the smallest value is", []int{3, 1, 4, 1, -10, 9, 2, 6}, -10, false},
		{"the smallest value is", []int{}, 0, true},
		{"the smallest value is", []int{3, 1, 4, 1, 5, 9, 2, 6}, 1, false},
	}
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := Min(tc.in)
            if tc.wantErr {
                if !errors.Is(err, ErrEmptySlice) {
                    t.Errorf("attendait ErrEmptySlice, got %v", err)
                }
                return
            }
            if err != nil {
                t.Fatalf("erreur inattendue: %v", err)
            }
            if got != tc.want {
                t.Errorf("Min(%v) = %d, want %d", tc.in, got, tc.want)
            }
        })
	}
}