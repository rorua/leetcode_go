package running_sum_of_1d_array

func runningSum(nums []int) []int {
	var (
		res = make([]int, len(nums))
		t   = 0
	)
	for k, v := range nums {
		res[k] = v + t
		t += v
	}
	return res
}
