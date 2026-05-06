package exo

import (
    "reflect"
    "testing"
)

func TestWordCount(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want map[string]int
    }{
        {"vide", "", map[string]int{}},
        {"un mot", "bonjour", map[string]int{"bonjour": 1}},
        {"répétition", "le chat le chien le chat", map[string]int{"le": 3, "chat": 2, "chien": 1}},
        {"casse mixte", "Le le LE", map[string]int{"le": 3}},
        {"espaces multiples", "  bonjour   monde  ", map[string]int{"bonjour": 1, "monde": 1}},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := WordCount(tc.in)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("WordCount(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}