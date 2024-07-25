package sort_the_jumbled_numbers

import (
	"fmt"
	"strconv"
)

func sortJumbled(mapping []int, nums []int) []int {

	if len(mapping) != 10 {
		return nil
	}

	var m = make([][2]int, len(nums))

	for i := 0; i < len(nums); i++ {
		m[i] = [2]int{
			nums[i],
			getMappedValue(nums[i], mapping),
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if m[j][1] < m[i][1] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}

	var res = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, m[i][0])
	}

	return res
}

func getMappedValue(num int, mapping []int) int {
	mappedValue := 0
	place := 1

	if num == 0 {
		return mapping[0]
	}

	for num > 0 {
		digit := num % 10
		mappedDigit := mapping[digit]
		mappedValue += mappedDigit * place
		place *= 10
		num /= 10
	}

	return mappedValue
}

func getMappedValue2(num int, mapping []int) int {
	strNum := fmt.Sprintf("%d", num)
	mappedStr := ""

	for _, char := range strNum {
		digit := char - '0'
		mappedDigit := mapping[digit]
		mappedStr += fmt.Sprintf("%d", mappedDigit)
	}

	mappedValue, _ := strconv.Atoi(mappedStr)
	return mappedValue
}
