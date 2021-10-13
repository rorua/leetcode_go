package sort_an_array

func sortArray(nums []int) []int {
	qsort(nums, 0, len(nums)-1)
	return nums
}

func getPivot(nums []int, pivot, end int) int {
	var swapIndex = pivot
	for i := pivot + 1; i <= end; i++ {
		if nums[i] < nums[pivot] {
			swapIndex++
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
		}
	}
	nums[pivot], nums[swapIndex] = nums[swapIndex], nums[pivot]
	return swapIndex
}

func qsort(nums []int, left, right int) {
	if left < right {
		pivot := getPivot(nums, left, right)
		qsort(nums, left, pivot-1)
		qsort(nums, pivot+1, right)
	}
}
