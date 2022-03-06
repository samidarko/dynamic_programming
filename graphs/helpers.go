package graphs

import "fmt"

// Queue data structure
type Queue []interface{}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return "", fmt.Errorf("empty queue")
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, nil
}

// Stack data structure
type Stack []interface{}

// Empty returns true if stack is empty
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Push a new element to the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

// Pop returns and remove the element on top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return "", fmt.Errorf("empty Stack")
	}
	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]
	return element, nil
}

// BuildGraph takes edges list and returns an adjacency list
func BuildGraph(edges [][]string) (map[string][]string, error) {
	g := map[string][]string{}
	for i, edge := range edges {
		if len(edge) != 2 {
			return g, fmt.Errorf("bad edge %v at index %d", edge, i)
		}
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	return g, nil
}
