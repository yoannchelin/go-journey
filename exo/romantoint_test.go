package exo

import (
    "errors"
    "testing"
)

func TestRomanToInt(t *testing.T) {
    tests := []struct {
        name    string
        in      string
        want    int
        wantErr bool
    }{
        {"III", "III", 3, false},
        {"IV", "IV", 4, false},
        {"IX", "IX", 9, false},
        {"LVIII", "LVIII", 58, false},
        {"MCMXCIV", "MCMXCIV", 1994, false},
        {"MMXXIV", "MMXXIV", 2024, false},
        {"vide", "", 0, true},
        {"chars invalides", "ABC", 0, true},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := RomanToInt(tc.in)
            if tc.wantErr {
                if !errors.Is(err, ErrInvalidRoman) {
                    t.Errorf("attendait ErrInvalidRoman, got %v", err)
                }
                return
            }
            if err != nil {
                t.Fatalf("erreur inattendue: %v", err)
            }
            if got != tc.want {
                t.Errorf("RomanToInt(%q) = %d, want %d", tc.in, got, tc.want)
            }
        })
    }
}