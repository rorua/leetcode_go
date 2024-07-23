package divide_array_into_equal_pairs

func divideArray(nums []int) bool {

	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}

	return true
}
