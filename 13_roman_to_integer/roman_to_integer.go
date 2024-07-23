package roman_to_integer

func romanToInt(s string) int {
	var romans = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res = romans[s[0]]
	var prev = romans[s[0]]

	for i := 1; i < len(s); i++ {
		cur := romans[s[i]]
		if prev < cur {
			res += cur - 2*prev
		} else {
			res += cur
		}
		prev = cur
	}

	return res
}
