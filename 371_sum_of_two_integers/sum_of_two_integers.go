package sum_of_two_integers

func getSum(a int, b int) int {
	var carry int
	for b != 0 {
		// & gives common set bits between a and b
		carry = a & b
		// ^ gives all set bits between a and b
		a = a ^ b
		// rightshift carry by 1 bit till we reach 0
		b = carry << 1
	}
	return a
}
