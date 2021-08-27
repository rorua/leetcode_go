package first_unique_character_in_a_string

//hash map solution
func firstUniqChar(s string) int {
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}

//array solution
func firstUniqChar2(s string) int {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i] - 'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i] - 'a'] == 1 {
			return i
		}
	}
	return -1
}
