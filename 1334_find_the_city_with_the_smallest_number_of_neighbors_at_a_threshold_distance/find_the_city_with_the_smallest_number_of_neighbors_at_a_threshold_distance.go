package find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance

//1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance
//
//There are n cities numbered from 0 to n-1. Given the array edges where edges[i] = [fromi, toi, weighti] represents
//a bidirectional and weighted edge between cities fromi and toi, and given the integer distanceThreshold.
//
//Return the city with the smallest number of cities that are reachable through some path and whose distance is at most
//distanceThreshold, If there are multiple such cities, return the city with the greatest number.
//
//Notice that the distance of a path connecting cities i and j is equal to the sum of the edges' weights along that
//path.

const inf int = 1e7

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	// Dijkstra's algorithm
	//var m = make(map[int]map[int]int, 0)
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := range m[i] {
			m[i][j] = inf
		}
	}

	//fmt.Printf("m=%v\n", m)

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		m[from][to], m[to][from] = weight, weight
	}

	//fmt.Printf("m=%v\n", m)

	var res, minCount = n, inf
	for i := n - 1; i >= 0; i-- {
		//count := dijkstra(i, n, m, distanceThreshold)
		//if count <= minCount {
		//	res, minCount = i, count
		//}
		if t := dijkstra(i, n, m, distanceThreshold); t < minCount {
			minCount, res = t, i
		}
	}

	return res
}

func dijkstra(source, n int, m [][]int, distanceThreshold int) (count int) {
	// Dijkstra's algorithm
	vis, dist := make([]bool, n), make([]int, n)
	for i := range vis {
		vis[i] = false
		dist[i] = inf
	}

	dist[source] = 0
	for i := 0; i < n; i++ {
		k := -1
		for j := 0; j < n; j++ {
			if !vis[j] && (k == -1 || dist[j] < dist[k]) {
				k = j
			}
		}

		vis[k] = true
		for j := 0; j < n; j++ {
			dist[j] = min(dist[j], dist[k]+m[k][j])
		}
	}

	for _, d := range dist {
		if d <= distanceThreshold {
			count++
		}
	}

	return
}
