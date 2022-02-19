package graphs

import "fmt"

func DepthFirstIterative(graph map[string][]string, source string) {
	s := Stack{source}

	for !s.Empty() {
		current, _ := s.Pop()
		fmt.Println(current)

		for _, neighbor := range graph[current.(string)] {
			s.Push(neighbor)
		}
	}
}

func DepthFirstRecursive(graph map[string][]string, source string) {
	fmt.Println(source)

	for _, neighbor := range graph[source] {
		DepthFirstRecursive(graph, neighbor)
	}
}

func BreadthFirstIterative(graph map[string][]string, source string) {
	q := Queue{source}

	for !q.Empty() {
		current, _ := q.Dequeue()
		fmt.Println(current)

		for _, neighbor := range graph[current.(string)] {
			q.Enqueue(neighbor)
		}
	}
}
