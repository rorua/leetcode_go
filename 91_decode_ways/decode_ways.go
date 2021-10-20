package decode_ways

func numDecodings(s string) int {
	res := make([]int, len(s)+1)
	res[0] = 1
	if s[0] == '0' {
		res[1] = 0
	} else {
		res[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			res[i] = res[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			res[i] += res[i-2]
		}
	}

	return res[len(s)]
}
