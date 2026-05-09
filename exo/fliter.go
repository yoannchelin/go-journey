package exo

func Filter(nums []int, pred func(int) bool) []int {
	result := []int{}

	for _, v := range nums[:] {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}