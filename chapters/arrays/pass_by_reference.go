package main

import "fmt"

func functionThatModifiesMatrix(matrix [][]int) {
	matrix[0][0] = 999
}

func main() {
	originalMatrix := [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("original matrix", originalMatrix)

	functionThatModifiesMatrix(originalMatrix)

	fmt.Println("modified matrix", originalMatrix)
}
