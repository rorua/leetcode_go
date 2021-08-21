package reverse_integer

func reverse(x int) int {

	if (x > 1<<31-1) || (x < -1<<31) {
		return 0
	}

	var res = 0
	for x != 0 {
		last := x % 10
		res = (res * 10) + last
		x = x / 10
	}

	if (res > 1<<31-1) || (res < -1<<31) {
		return 0
	}

	return res
}
