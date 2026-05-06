package warmup2

import "errors"

var ErrEmptySlice = errors.New("SLICE VIDE")

func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return nums[0], ErrEmptySlice
	}

	maxi := nums[0]
	for _, i := range nums[1:] {
		if i > maxi {
			maxi = i
		}
	}
	return maxi, nil
}