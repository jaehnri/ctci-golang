package main

import "fmt"

const (
	UNCHECKED = 0
	CHECKING  = 1
	SAFE      = 2
)

// IsNodeSafe returns whether a node leads to no loops. A node is safe
// iff it only leads to other nodes that are safe.
func IsNodeSafe(edges [][]int, statuses []int, current int) bool {
	statuses[current] = CHECKING
	for _, neighbor := range edges[current] {
		if statuses[neighbor] == CHECKING {
			return false
		}

		if statuses[neighbor] == UNCHECKED {
			safe := IsNodeSafe(edges, statuses, neighbor)
			if !safe {
				return false
			}
		}
	}

	statuses[current] = SAFE
	return true
}

func CycleInGraph(edges [][]int) bool {
	statuses := make([]int, len(edges))

	for index := range edges {
		if statuses[index] == UNCHECKED {
			safe := IsNodeSafe(edges, statuses, index)
			if !safe {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(CycleInGraph([][]int{
		{1, 3},
		{2, 3, 4},
		{0},
		{},
		{2, 5},
		{},
	}))
}
