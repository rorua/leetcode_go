package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var dict = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res = make([]string, 0)
	res = append(res, "")

	for _, digit := range digits {
		var tmp []string
		for _, str := range res {
			for _, c := range dict[byte(digit)] {
				tmp = append(tmp, str+string(c))
			}
		}
		res = tmp
	}

	return res
}
