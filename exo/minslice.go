package exo

import "errors"

var ErrorEmptySlice = errors.New("SLICE EMPTY")

func Min(nums []int) (int, error) {
	if len(nums)-1 < 0 {
		return 0, ErrorEmptySlice
	}
	min := nums[0]
	for _, i := range nums[1:] {
		if i < min {
			min = i
		}
	}
	return min, nil
}