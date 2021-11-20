package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[l] > target || target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[m] > target || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
