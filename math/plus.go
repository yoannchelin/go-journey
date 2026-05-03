package math

import (
	"errors"
	"fmt"
)

func Plus(a, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

func MainPlus(a, b, c any) error {
	av, ok1 := a.(int)
	bv, ok2 := b.(int)
	cv, ok3 := c.(int)

	if !ok1 || !ok2 || !ok3 {
		return errors.New("valeur numérique invalide")
	}

	res := Plus(av, bv)
	fmt.Printf("le resultat de l'operation est %d\n", res)
	zeb := PlusPlus(av, bv, cv)
	fmt.Printf("le resultat de l'operation est %d\n", zeb)

	return nil
}