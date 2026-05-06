package main

import (
	"example/user/hello/exo"
	"fmt"
)

func main() {
	fmt.Print(exo.ReverseString(""))
	fmt.Print(exo.Min([]int{3, 1, 4, 1, 5, 9, 2, 6}))
	fmt.Print(exo.CountVowels("je sais pas quoi dire mon vierrrr"))
}
