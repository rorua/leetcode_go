package climbing_stairs

func climbStairs(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

func climbStairs2(n int) int {
	one, two := 1, 1
	for i := 0; i < n; i++ {
		one, two = one+two, one
	}
	return one
}
