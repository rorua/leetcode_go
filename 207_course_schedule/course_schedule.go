package course_schedule

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
