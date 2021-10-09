package single_number

func singleNumber(nums []int) int {
	uniq := 0
	for _, num := range nums {
		uniq ^= num
	}
	return uniq
}

func singleNumber2(nums []int) int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
