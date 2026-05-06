package exo

import "errors"

var ErrEmptySlice = errors.New("slice vide")

func Min(nums []int) (int, error) {

	if len(nums)-1 == 0 {
		return 0, ErrEmptySlice
	}
	min := nums[0]
	for _, i := range nums[:] {
		if i < min {
			min = i
		}
	}
	return min, nil
}