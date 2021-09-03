package contains_duplicate

func containsDuplicate(nums []int) bool {
	var m = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			return true
		}
		m[nums[i]]++
	}
	return false
}
