package house_robber

func rob(nums []int) int {
	rob1, rob2 := 0, 0
	// [rob1, rob2, n, n+1, ...]
	for _, n := range nums {
		tmp := func() int {
			if n+rob1 > rob2 {
				return n + rob1
			}
			return rob2
		}()
		rob1 = rob2
		rob2 = tmp
	}
	return rob2
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := make([]int, len(nums)+1)
	res[0] = 0
	res[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i+1] = func() int {
			if res[i] > res[i-1]+nums[i] {
				return res[i]
			}
			return res[i-1] + nums[i]
		}()
	}
	return res[len(nums)]
}
