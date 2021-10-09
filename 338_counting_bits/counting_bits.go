package counting_bits

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = func(k int) int {
			cnt := 0
			for k > 0 {
				cnt += k & 1
				k >>= 1
			}
			return cnt
		}(i)
	}
	return res
}
