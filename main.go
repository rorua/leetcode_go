package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	visited := make(map[int]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if visit(graph, visited, i) {
			return false
		}
	}
	return true
}

func visit(graph [][]int, visited map[int]int, vertex int) bool {
	visited[vertex]++
	for _, neighbour := range graph[vertex] {
		if visited[neighbour] > 0 || visit(graph, visited, neighbour) {
			return true
		}
	}
	delete(visited, vertex)
	return false
}

func singleNumber(nums []int) int {
	uniq := 0
	for _, num := range nums {
		curU := uniq
		uniq = uniq ^ num
		fmt.Printf("cur = %d, num = %d, res=%d^%d = %d\n", curU, num, curU, num, uniq)
	}
	return uniq
}

func moveZeroes(nums []int) {
	l, r := 0, 0
	for l < len(nums) && r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func main() {
	moveZeroes([]int{1, 2, 0, 3, 0, 4, 3, 0, 5})
}
