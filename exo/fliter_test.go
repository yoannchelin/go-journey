package exo

import (
    "testing"
    "reflect"
)

func TestFilter(t *testing.T) {
    isEven := func(n int) bool { return n%2 == 0 }
    isPositive := func(n int) bool { return n > 0 }

    tests := []struct {
        name string
        in []int
        pred func(int) bool
        want []int
    } {
        {"slice normal", []int{1, 2, 3, 4, 5}, isEven, []int{2, 4}},
        {"slice vide", []int{}, isPositive, []int{}},
        {"slice valeur neg", []int{-10, -20, 3, 4, 5}, isEven, []int{-10, -20 ,4}},
        {"slice valeur neg positiv", []int{-10, -20, 3, 4, 5}, isPositive, []int{3, 4 ,5}},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Filter(tc.in, tc.pred)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("Filter(%v) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}