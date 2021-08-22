package buddy_strings

func buddyStrings(s string, goal string) bool {

	var n = len(s)
	if n != len(goal) {
		return false
	}

	var (
		chars = make([]byte, 0)
		m     = make(map[byte]int)
	)

	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			chars = append(chars, s[i], goal[i])
		}
		m[s[i]]++
	}

	if s == goal && len(m) < n {
		return true
	}

	if len(chars) == 4 && chars[0] == chars[3] && chars[1] == chars[2] {
		return true
	}

	return false
}
