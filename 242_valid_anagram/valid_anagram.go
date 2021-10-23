package valid_anagram

func isAnagram(s string, target string) bool {
	chars := make(map[rune]int)
	for _, v := range s {
		chars[v]++
	}

	for _, v := range target {
		chars[v]--
	}

	for k, _ := range chars {
		if chars[k] != 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, target string) bool {
	if len(s) != len(target) {
		return false
	}

	chars := make(map[string]int)
	for i := 0; i < len(s); i++ {
		chars[string(s[i])]++
		chars[string(target[i])]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
