package sort_the_people

func sortPeople(names []string, heights []int) []string {

	var (
		n = len(heights)
		m = make(map[int]string, n)
	)

	for i := 0; i < n; i++ {
		m[heights[i]] = names[i]
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if heights[j] > heights[i] {
				heights[i], heights[j] = heights[j], heights[i]
			}
		}
	}

	var res = make([]string, 0)
	for i := 0; i < len(heights); i++ {
		res = append(res, m[heights[i]])
	}

	return res
}

func sortPeople2(names []string, heights []int) []string {
	n := len(heights)
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	// Sort indices based on heights
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if heights[indices[j]] > heights[indices[i]] {
				indices[i], indices[j] = indices[j], indices[i]
			}
		}
	}

	res := make([]string, n)
	for i, idx := range indices {
		res[i] = names[idx]
	}

	return res
}
