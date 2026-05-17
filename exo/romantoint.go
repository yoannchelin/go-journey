package exo

import (
	"errors"
)

var ErrInvalidRoman = errors.New("Invalid roman")

func RomanToInt(s string) (int, error) {
	values := map[rune]int{
    	'I': 1,
    	'V': 5,
    	'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if s == "" {
    	return 0, ErrInvalidRoman
	}
	total := 0

	for i, r := range s {
		// Étape 1 : récupère la valeur de r dans la map
		v, ok := values[r]
		if !ok {
			// caractère invalide
			return 0, ErrInvalidRoman
		}

		// Étape 2 : est-ce qu'il y a un caractère suivant ?
		if i+1 < len(s) {
			// Étape 3 : récupère la valeur du suivant
			next := values[rune(s[i+1])]
			if v < next {
				// soustraire
				total -= v
				continue
			}
		}

		// Étape 4 : additionner
		total += v
	}

	return total, nil
}