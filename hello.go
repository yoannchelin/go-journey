package main

import (
	"example/user/hello/warmup2"
	"fmt"
)

func main() {
	even := warmup2.IsEven(3)
	if even == false {
		fmt.Printf("c'est impair")
	}
	if even == true {
		fmt.Printf("c'est pair")
	}
}
