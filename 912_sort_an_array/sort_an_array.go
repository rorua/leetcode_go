package sort_an_array

// 912. Sort an Array
//
// Given an array of integers nums, sort the array in ascending order and return it.
//
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the
// smallest space complexity possible.
//
//Constraints:
//
// 1 <= nums.length <= 5 * 10^4
//-5 * 10^4 <= nums[i] <= 5 * 10^4

func sortArray(nums []int) []int {
	//qsort(nums, 0, len(nums)-1)
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func getPivot(nums []int, pivot, end int) int {
	var swapIndex = pivot
	for i := pivot + 1; i <= end; i++ {
		if nums[i] < nums[pivot] {
			swapIndex++
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
		}
	}
	nums[pivot], nums[swapIndex] = nums[swapIndex], nums[pivot]
	return swapIndex
}

func qsort(nums []int, left, right int) {
	if left < right {
		pivot := getPivot(nums, left, right)
		qsort(nums, left, pivot-1)
		qsort(nums, pivot+1, right)
	}
}

func mergeSort(nums []int, left, right int) {
	if left == right {
		return
	}

	//fmt.Printf("left: %d, right: %d, nums=%v\n", left, right, nums)

	mid := left + (right-left)/2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)

	merge(nums, left, mid, right)
}

func merge(nums []int, left int, mid int, right int) {
	//lArr, rArr := nums[left:mid+1], nums[mid+1:right+1]
	lArr, rArr := make([]int, mid-left+1), make([]int, right-mid)
	copy(lArr, nums[left:mid+1])
	copy(rArr, nums[mid+1:right+1])
	i, j, k := left, 0, 0
	for j < len(lArr) && k < len(rArr) {
		if lArr[j] <= rArr[k] {
			nums[i] = lArr[j]
			j++
		} else {
			nums[i] = rArr[k]
			k++
		}
		i++
	}

	for j < len(lArr) {
		nums[i] = lArr[j]
		j++
		i++
	}

	for k < len(rArr) {
		nums[i] = rArr[k]
		k++
		i++
	}
}
