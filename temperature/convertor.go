package temperature

import (
	"errors"
	"strconv"
)

const (
	celsiusToFahrenheitFactor = 9.0 / 5.0
	fahrenheitToCelsiusFactor = 5.0 / 9.0
	fahrenheitOffset          = 32.0
)

func ConvertorTemp(x string) (string, error) {
	if len(x) < 2 {
		return "", errors.New("input trop court")
	}

	lastChar := x[len(x)-1]
	numStr := x[:len(x)-1]

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return "", errors.New("valeur numérique invalide")
	}

	switch lastChar {
	case 'C':
		result := celsiusToFahrenheitFactor*num + fahrenheitOffset
		return strconv.FormatFloat(result, 'f', 2, 64) + "F", nil
	case 'F':
		result := fahrenheitToCelsiusFactor * (num - fahrenheitOffset)
		return strconv.FormatFloat(result, 'f', 2, 64) + "C", nil
	}

	return "", errors.New("unité inconnue, utilise C ou F")
}
