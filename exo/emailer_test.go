package exo

import "testing"

func TestIsValidEmail(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want bool
    }{
        {"valide simple", "yoann@gmail.com", true},
        {"valide complexe", "first.last+tag@domain.co.uk", true},
        {"vide", "", false},
        {"sans @", "yoanngmail.com", false},
        {"deux @", "yoann@@gmail.com", false},
        {"@ au début", "@gmail.com", false},
        {"@ à la fin", "yoann@", false},
        {"pas de point dans domain", "yoann@gmail", false},
        {"point juste après @", "yoann@.com", false},
        {"avec espace", "yo ann@gmail.com", false},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := IsValidEmail(tc.in)
            if got != tc.want {
                t.Errorf("IsValidEmail(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}