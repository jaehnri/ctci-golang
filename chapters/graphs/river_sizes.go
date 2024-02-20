package main

func RiverSizes(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	riverLengths := []int{}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				continue
			}

			if matrix[i][j] == 1 {
				riverLengths = append(riverLengths, calculateRiverSize(matrix, visited, i, j))
			}
		}
	}

	return riverLengths
}

func calculateRiverSize(matrix [][]int, visited [][]bool, i, j int) int {
	up, down, left, right := 0, 0, 0, 0
	visited[i][j] = true

	// check up
	if i > 0 && !visited[i-1][j] && matrix[i-1][j] == 1 {
		up = calculateRiverSize(matrix, visited, i-1, j)
	}

	// check left
	if j > 0 && !visited[i][j-1] && matrix[i][j-1] == 1 {
		left = calculateRiverSize(matrix, visited, i, j-1)
	}

	// check down
	if i < len(matrix)-1 && !visited[i+1][j] && matrix[i+1][j] == 1 {
		down = calculateRiverSize(matrix, visited, i+1, j)
	}

	// check right
	if j < len(matrix[0])-1 && !visited[i][j+1] && matrix[i][j+1] == 1 {
		right = calculateRiverSize(matrix, visited, i, j+1)
	}

	return 1 + up + down + left + right
}

//func main() {
//	fmt.Println(RiverSizes([][]int{
//		{1, 0, 0, 1, 0},
//		{1, 0, 1, 0, 0},
//		{0, 0, 1, 0, 1},
//		{1, 0, 1, 0, 1},
//		{1, 0, 1, 1, 0},
//	}))
//}
