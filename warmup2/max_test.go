package warmup2

import (
	"errors"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		in int
		want int
	} {
		{"la valeur max est %d", 5,5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {}) {
			got := Max(tc.in)
			if got != Max(tc.want) {
				t.Errorf("la slice est vide")
			}
		}
	}
}