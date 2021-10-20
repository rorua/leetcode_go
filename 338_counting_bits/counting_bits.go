package counting_bits

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = func(k int) int {
			count := 0
			for k > 0 {
				count += k & 1
				k >>= 1
			}
			return count
		}(i)
	}
	return res
}
