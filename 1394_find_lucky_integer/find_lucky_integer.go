package find_lucky_integer

func findLucky(arr []int) int {
	var nums = make(map[int]int)
	for _, num := range arr {
		nums[num]++
	}
	res := -1
	for k, v := range nums {
		if k == v && v > res {
			res = v
		}
	}
	return res
}
