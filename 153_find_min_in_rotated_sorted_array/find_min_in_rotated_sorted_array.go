package find_min_in_rotated_sorted_array

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	m := 0
	for nums[l] > nums[r] {
		m = (l + r) / 2
		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}

func findMin2(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] <= nums[r] {
			res = min(res, nums[l])
			break
		}
		m := (l + r) / 2
		res = min(res, nums[m])
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
