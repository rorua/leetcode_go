package check_if_one_string_swap_can_make_strings_equal

func areAlmostEqual(s1 string, s2 string) bool {
	var n = len(s1)
	if s1 == s2 {
		return true
	}

	var (
		m = make(map[byte]int)
		k = 0
	)

	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			k++
		}
		if k > 2 {
			return false
		}
		m[s1[i]]++
		m[s2[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
