package graphs

func ShortestPath(edges [][]string, startNode, endNode string) int {
	graph, _ := BuildGraph(edges)
	visited := map[string]bool{startNode: true}
	type Element struct {
		node     string
		distance int
	}
	queue := []Element{{startNode, 0}}

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		if element.node == endNode {
			return element.distance
		}

		for _, neighbor := range graph[element.node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, Element{neighbor, element.distance + 1})
			}
		}
	}
	return -1
}
