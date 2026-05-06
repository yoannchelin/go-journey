package exo

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in string
		want string
	} {
		{"reverse string is ", "a", "a"},
		{"reverse string is ", "", ""},
		{"reverse string is ", "hello", "olleh"},
		{"reverse string is ", "héllo", "olléh"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T)  {
			got := ReverseString(tc.in)
			if got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}