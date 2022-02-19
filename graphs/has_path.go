package graphs

// HasPathRecursive check if graph has path between source and target recursively
func HasPathRecursive(graph map[string][]string, source, target string) bool {
	if source == target {
		return true
	}

	for _, neighbor := range graph[source] {
		if HasPathRecursive(graph, neighbor, target) {
			return true
		}
	}

	return false
}

// HasPathIterative check if graph has path between source and target iteratively
func HasPathIterative(graph map[string][]string, source, target string) bool {
	q := Queue{source}

	for !q.Empty() {
		current, _ := q.Dequeue()
		if current == target {
			return true
		}

		for _, neighbor := range graph[current.(string)] {
			q.Enqueue(neighbor)
		}
	}

	return false
}

func HasPathUndirected(graph map[string][]string, source, target string, visited map[string]bool) bool {
	if source == target {
		return true
	}

	if visited[source] {
		return false
	}

	visited[source] = true

	for _, neighbor := range graph[source] {
		if HasPathUndirected(graph, neighbor, target, visited) {
			return true
		}
	}

	return false
}
