package main

func TransposeMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	// initialize transposed matrix
	transposed := make([][]int, n)
	for i := range transposed {
		transposed[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

//func main() {
//	fmt.Println(TransposeMatrix([][]int{{1, 2, 3}, {4, 5, 6}}))
//}
