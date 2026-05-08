package exo


func Filter(nums []int, pred func (int) bool) []int {

	res := []int{}
	for _, i := range nums {
		if pred(i) {
			res = append(res, i)
		}
	}
	return res
}