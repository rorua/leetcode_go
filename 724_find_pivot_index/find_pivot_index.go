package find_pivot_index

func pivotIndex(nums []int) int {
	sum := 0
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if sum == total-sum-nums[i] {
			return i
		}
		sum += nums[i]
	}

	return -1
}
