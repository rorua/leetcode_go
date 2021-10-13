package maximum_subarray

func maxSubArray(nums []int) int {
	max := nums[0]
	curSum := 0
	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		max = func(max, curSum int) int {
			if max > curSum {
				return max
			}
			return curSum
		}(max, curSum)
	}
	return max
}
