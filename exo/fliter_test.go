package exo



import (
    "reflect"
    "testing"
)

func TestFilter(t *testing.T) {
    isEven := func(n int) bool { return n%2 == 0 }
    isPositive := func(n int) bool { return n > 0 }
    isMagic := func(n int) bool { return n == 42 }

    tests := []struct {
        name string
        in   []int
        pred func(int) bool
        want []int
    }{
        {"pairs", []int{1, 2, 3, 4, 5}, isEven, []int{2, 4}},
        {"positifs", []int{-2, -1, 0, 1, 2}, isPositive, []int{1, 2}},
        {"slice vide", []int{}, isEven, []int{}},
        {"aucun match", []int{1, 3, 5}, isEven, []int{}},
        {"que des matches", []int{42, 42, 42}, isMagic, []int{42, 42, 42}},
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