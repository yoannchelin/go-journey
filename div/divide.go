package div

import (
	"errors"

)



func Divide(a, b int) (int, int, error) {


		if b == 0 {
		return a, b, errors.New("pas divisible par 0")
	}

	c := a / b

	return c, a%b, nil
}