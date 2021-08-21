package longest_substring

func lengthOfLongestSubstring(s string) int {
	var n = len(s)
	if n == 0 || n == 1 {
		return n
	}
	var (
		substr       = make(map[byte]int)
		i, j, max, l = 0, 1, 0, 0
	)
	substr[s[0]]++
	for j < n {
		c := s[j]
		substr[c]++
		for substr[c] == 2 && i < j {
			substr[s[i]]--
			if substr[s[i]] == 0 {
				delete(substr, s[i])
			}
			i++
		}
		l = len(substr)
		if max < l {
			max = l
		}
		j++
	}
	return max
}
