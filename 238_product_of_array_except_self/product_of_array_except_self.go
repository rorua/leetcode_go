package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	postfix := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			prefix[i] = nums[i]
			postfix[n-i-1] = nums[n-i-1]
		} else {
			prefix[i] = nums[i] * prefix[i-1]
			postfix[n-i-1] = nums[n-i-1] * postfix[n-i]
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			nums[i] = postfix[i+1]
		} else if i != n-1 {
			nums[i] = prefix[i-1] * postfix[i+1]
		} else {
			nums[i] = prefix[i-1]
		}

	}

	return nums
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
