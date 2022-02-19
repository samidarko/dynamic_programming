package graphs

import (
	"testing"
)

func TestDepthFirstIterative(t *testing.T) {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	DepthFirstIterative(graph, "a")
}

func TestDepthFirstRecursive(t *testing.T) {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	DepthFirstRecursive(graph, "a")
}

func TestBreadthFirstIterative(t *testing.T) {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	BreadthFirstIterative(graph, "a")
}
