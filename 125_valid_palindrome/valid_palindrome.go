package valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlphaNum(s[l]) {
			l++
		}
		for l < r && !isAlphaNum(s[r]) {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}
