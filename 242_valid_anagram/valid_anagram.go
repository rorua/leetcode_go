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
