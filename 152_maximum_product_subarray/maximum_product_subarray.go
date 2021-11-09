package maximum_product_subarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	curMin, curMax := 1, 1

	for _, n := range nums {
		curMax, curMin = max(n*curMax, n*curMin, n), min(n*curMax, n*curMin, n)
		res = max(res, curMax)
	}

	return res
}

func min(vals ...int) int {
	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

func max(vals ...int) int {
	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}
