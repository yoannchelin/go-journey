package exo

import (
	"errors"
	"testing"
)

func TestMin(t *testing.T) {
    tests := []struct {
        name string
        in []int
        want int
        wantErr bool
    } {
        {"slice vide", []int{}, 0, true},
        {"slice nbre neg", []int{1,2,3,4,-10}, -10, false},
        {"slice normal", []int{1,2,3,4,5}, 0, false},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := Min(tc.in)
            if tc.wantErr {
                if !errors.Is(ErrorEmptySlice, err) {
                    t.Errorf("waiting for ErrEmptyslice got %v", err)
                }
                return
            }
            if err != nil {
                t.Errorf("erreur innatendue %v", err)
            }
            if got != tc.want {
                t.Errorf("Min(%d) = %d, want %d", tc.in, got, tc.want)
            }
        })
    }
}