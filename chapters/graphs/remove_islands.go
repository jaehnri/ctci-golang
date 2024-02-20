package main

import "fmt"

const Mainland = 2

func RemoveIslands(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, 0

	// check up borders
	for j < n {
		if matrix[i][j] == 1 {
			markAsMainland(matrix, i, j)
		}
		j++
	}

	// check right borders
	j--
	for i < m {
		if matrix[i][j] == 1 {
			markAsMainland(matrix, i, j)
		}
		i++
	}

	// check down borders
	i--
	for j >= 0 {
		if matrix[i][j] == 1 {
			markAsMainland(matrix, i, j)
		}
		j--
	}

	// check left borders
	j++
	for i >= 0 {
		if matrix[i][j] == 1 {
			markAsMainland(matrix, i, j)
		}
		i--
	}

	i = 0
	j = 0
	for i = range matrix {
		for j = range matrix[i] {
			if matrix[i][j] != Mainland {
				matrix[i][j] = 0
			}

			if matrix[i][j] == Mainland {
				matrix[i][j] = 1
			}
		}
	}

	return matrix
}

func markAsMainland(matrix [][]int, i, j int) {
	matrix[i][j] = Mainland

	// check up
	if i > 0 && matrix[i-1][j] == 1 {
		markAsMainland(matrix, i-1, j)
	}

	// check down
	if i < len(matrix)-1 && matrix[i+1][j] == 1 {
		markAsMainland(matrix, i+1, j)
	}

	// check left
	if j > 0 && matrix[i][j-1] == 1 {
		markAsMainland(matrix, i, j-1)
	}

	// check right
	if j < len(matrix[0])-1 && matrix[i][j+1] == 1 {
		markAsMainland(matrix, i, j+1)
	}
}

func main() {
	fmt.Println(RemoveIslands(
		[][]int{
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 1, 1, 1},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 0, 1, 1, 0, 0},
			{1, 0, 0, 0, 0, 1},
		},
	))
}
