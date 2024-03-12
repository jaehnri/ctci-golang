package main

import "math"

const (
	UncheckedCell = -2
	OnCheck       = -1
)

func MinimumPassesOfMatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	visited := initializeVisitedMatrix(m, n)
	maxPasses := math.MaxInt

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				// check up
				if i > 0 && matrix[i-1][j] > 0 {
					visited[i][j] = 1
					continue
				}

				// check down
				if i < m-1 && matrix[i+1][j] > 0 {
					visited[i][j] = 1
					continue
				}

				// check left
				if j > 0 && matrix[i][j-1] > 0 {
					visited[i][j] = 1
					continue
				}

				// check right
				if j < n-1 && matrix[i][j+1] > 0 {
					visited[i][j] = 1
					continue
				}

			}
		}
	}

	return maxPasses
}

//func FindMin(matrix, visited [][]int, i, j int) int {
//	//if matrix[i][j] == 0 {
//	//	visited[i][j] = 0
//	//	return 0
//	//}
//	//
//	//if matrix[i][j] > 0 {
//	//		[visite]
//	//}
//
//	if i > 0 && matrix[i-1][j] > 0 {
//		visited[i][j] = 1
//		continue
//	}
//
//	// check down
//	if i < m-1 && matrix[i+1][j] > 0 {
//		visited[i][j] = 1
//		continue
//	}
//
//	// check left
//	if j > 0 && matrix[i][j-1] > 0 {
//		visited[i][j] = 1
//		continue
//	}
//
//	// check right
//	if j < n-1 && matrix[i][j+1] > 0 {
//		visited[i][j] = 1
//		continue
//	}
//
//}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initializeVisitedMatrix(m, n int) [][]int {
	visited := make([][]int, m)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = UncheckedCell
		}
		i++
	}
	return visited
}

func main() {

}
