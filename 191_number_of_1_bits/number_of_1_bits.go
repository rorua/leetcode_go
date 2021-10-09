package number_of_1_bits

func hammingWeight(num uint32) int {
	cnt := 0
	for num > 0 {
		cnt += int(num) & 1
		num >>= 1
	}
	return cnt
}
