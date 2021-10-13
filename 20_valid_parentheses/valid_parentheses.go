package valid_parentheses

func isValid(s string) bool {
	m := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	stack := make([]string, 0)
	for _, c := range s {
		if v, found := m[string(c)]; found {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(c))
		}
	}
	return len(stack) == 0
}

func IsValid(s string) bool {
	return isValid(s)
}
