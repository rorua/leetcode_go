package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var (
		reversed = 0
		original = x
	)
	for x != 0 {
		last := x % 10
		reversed = (reversed * 10) + last
		x = x / 10
	}

	return original == reversed
}
