package rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[(i+k)%n] = nums[i]
	}
	copy(nums, arr)
}
