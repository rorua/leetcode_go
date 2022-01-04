package container_with_most_water

func maxArea(height []int) int {

	var (
		res = 0
		l   = 0
		r   = len(height) - 1
	)

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

//brute force
func maxArea2(height []int) int {

	var res = 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[j], height[i])
			res = max(res, area)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
