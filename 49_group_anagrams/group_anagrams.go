package group_anagrams

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	dict := make(map[[26]int][]string, 0)
	for _, str := range strs {
		gr := [26]int{}
		for _, c := range str {
			gr[c-'a']++
		}
		if _, found := dict[gr]; !found {
			dict[gr] = make([]string, 0)
		}
		dict[gr] = append(dict[gr], str)
	}

	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	res := make([][]string, 0)
	dict := make(map[[26]byte]int)
	for _, str := range strs {
		gr := [26]byte{}
		for c := range str {
			gr[str[c]-'a']++
		}
		if idx, found := dict[gr]; found {
			res[idx] = append(res[idx], str)
		} else {
			res = append(res, []string{str})
			dict[gr] = len(res) - 1
		}
	}
	return res
}

func groupAnagrams3(strs []string) [][]string {
	res := make([][]string, 0)
	dict := make(map[[26]byte]int)
	for i := range strs {
		gr := [26]byte{}
		for j := range strs[i] {
			gr[strs[i][j]-'a']++
		}
		if idx, found := dict[gr]; found {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			dict[gr] = len(res) - 1
		}
	}
	return res
}
