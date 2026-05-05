package warmup2

import "errors"


func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return nums[0], errors.New("slice vide")
	}

	maxi := nums[0]
	for i := range nums {
		maxi = i
	}
	return maxi, nil
}