package binary_search

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end && end < len(nums) {
		ind := start + (end-start)/2
		if nums[ind] == target {
			return ind
		}
		if target > nums[ind] {
			start = ind + 1
		} else {
			end = ind - 1
		}
	}
	return -1
}
