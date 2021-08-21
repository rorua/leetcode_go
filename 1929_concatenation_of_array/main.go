package main

import (
	"fmt"
	"strconv"
)

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
	strconv.Atoi("123")
}
