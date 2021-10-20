package missing_number

//using hashmap
func missingNumber(nums []int) int {
	m := make(map[int]int)
	for i := 1; i <= len(nums); i++ {
		m[nums[i-1]]++
	}

	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			return i
		}
	}
	return 0
}

//using xor
func missingNumber2(nums []int) int {
	res := 0
	for i := 0; i <= len(nums); i++ {
		res ^= i
	}
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

//using xor2
func missingNumber3(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	return res
}

//using sum
func missingNumber4(nums []int) int {
	sum1, sum2 := 0, 0
	for i := 0; i <= len(nums); i++ {
		sum1 += i
	}
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	return sum1 - sum2
}

//using sum2
func missingNumber5(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res += i - nums[i]
	}

	return res
}
