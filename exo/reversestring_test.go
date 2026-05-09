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
		{"la string reverse est", "abc", "cba"},
	  	{"la string reverse est", "", ""},
	  	{"la string reverse est", "zebi", "ibez"},
	}

	  for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseRunes(tc.in)
			if got != tc.want {
				t.Errorf("ReverseRunes(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	  }
}