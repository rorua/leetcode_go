package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		for _, str := range strs {
			if i == len(str) || str[i] != c {
				return prefix[:i]
			}
		}
	}
	return prefix
}
