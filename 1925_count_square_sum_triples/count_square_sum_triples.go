package count_square_sum_triples

func countTriples(n int) int {
	m := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		m[i*i] = i
	}
	res := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if v, found := m[i*i+j*j]; found {
				if v > j {
					res += 2
				}
			}
		}
	}
	return res
}
