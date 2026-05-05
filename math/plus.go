package math

import (
	"errors"
	"fmt"
)

// ErrInvalidNumber est retournée quand un argument n'est pas un int.
var ErrInvalidNumber = errors.New("valeur numérique invalide")

// Plus retourne la somme de deux entiers.
func Plus(a, b int) int {
	return a + b
}

// PlusPlus retourne la somme de trois entiers.
func PlusPlus(a, b, c int) int {
	return a + b + c
}

// toInt convertit une valeur any en int.
// Retourne ErrInvalidNumber si la valeur n'est pas un int.
func toInt(v any) (int, error) {
	n, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("%w: %v (type %T)", ErrInvalidNumber, v, v)
	}
	return n, nil
}

// SumAny additionne trois valeurs any en les convertissant en int.
// Retourne le résultat de Plus(a,b), de PlusPlus(a,b,c), ou une erreur.
func SumAny(a, b, c any) (sum2, sum3 int, err error) {
	av, err := toInt(a)
	if err != nil {
		return 0, 0, err
	}
	bv, err := toInt(b)
	if err != nil {
		return 0, 0, err
	}
	cv, err := toInt(c)
	if err != nil {
		return 0, 0, err
	}

	return Plus(av, bv), PlusPlus(av, bv, cv), nil
}
