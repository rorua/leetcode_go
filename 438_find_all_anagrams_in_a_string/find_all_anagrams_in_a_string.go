package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	n := len(p)
	res := make([]int, 0)
	for i := 0; i+n-1 < len(s); i++ {
		tmp := s[i : i+n]
		isAnagram := func(str, target string) bool {
			chars := make(map[rune]int)
			for _, v := range str {
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
		}(tmp, p)

		if isAnagram {
			res = append(res, i)
		}
	}

	return res
}
