package minimum_deletions_to_make_string_balanced

//You are given a string s consisting only of characters 'a' and 'b'.
//
//You can delete any number of characters in s to make s balanced. s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
//
//Return the minimum number of deletions needed to make s balanced.

//Constraints:
//
//1 <= s.length <= 10^5
//s[i] is 'a' or 'b'.

func minimumDeletions(s string) int {

	var (
		res         = len(s)
		aCountRight = make([]int, len(s))
		bCountLeft  = 0
	)

	for i := len(s) - 2; i >= 0; i-- {
		aCountRight[i] = aCountRight[i+1]
		if s[i+1] == 'a' {
			aCountRight[i]++
		}
	}

	for i := 0; i < len(s); i++ {
		res = min(res, bCountLeft+aCountRight[i])
		if s[i] == 'b' {
			bCountLeft++
		}
	}

	return res
}
