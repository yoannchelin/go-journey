package exo

import "testing"


func TestIsValidEmail(t *testing.T) {
    tests := []struct {
        name string
        in string
        want bool
    } {
        {"email vide", "", false},
        {"email valide", "yoann@gmail.com", true},
        {"email avec espace", "yoa nn@gmail.com", false},
        {"deux @", "yoann@@gmail.com", false},
        {"@ premier argument", "@gmail.com", false},
        {"email sans domain", "yoann@.com", false},
        {"email sans point", "yoann@gmail", false},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := IsValidEmail(tc.in)
            if got != tc.want {
                t.Errorf("IsEmailValid(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}