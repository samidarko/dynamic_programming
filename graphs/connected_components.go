package graphs

func ConnectedComponentsCount(graph map[int][]int) int {
	visited := map[int]bool{}
	count := 0
	for node, _ := range graph {
		if ExploreGraph(graph, node, visited) {
			count++
		}
	}
	return count
}

func ExploreGraph(graph map[int][]int, current int, visited map[int]bool) bool {
	if visited[current] {
		return false
	}

	visited[current] = true

	for _, neighbor := range graph[current] {
		ExploreGraph(graph, neighbor, visited)
	}

	return true
}

func LargestComponent(graph map[int][]int) int {
	visited := map[int]bool{}
	largest := 0
	for node, _ := range graph {
		count := countNodes(graph, node, visited)
		if count > largest {
			largest = count
		}

	}
	return largest
}

func countNodes(graph map[int][]int, current int, visited map[int]bool) int {
	if visited[current] {
		return 0
	}

	visited[current] = true
	count := 1

	for _, neighbor := range graph[current] {
		count += countNodes(graph, neighbor, visited)
	}

	return count
}
