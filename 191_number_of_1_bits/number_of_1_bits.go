package number_of_1_bits

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num) % 2
		num >>= 1
	}
	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num) & 1
		num >>= 1
	}
	return count
}

func hammingWeight3(num uint32) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count += 1
	}
	return count
}
