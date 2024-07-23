package sort_array_by_increasing_frequency

func frequencySort(nums []int) []int {

	var (
		n = len(nums)
		m = make(map[int]int, n)
	)

	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[nums[j]] < m[nums[i]] || (m[nums[j]] == m[nums[i]] && nums[j] > nums[i]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
