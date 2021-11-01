package search_insert_position

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		}else{
			r = mid - 1
		}
	}
	return l
}
