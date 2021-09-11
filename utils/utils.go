package utils

func InArray(val string, array []string) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func InArrayInt(val int, array []int) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}
