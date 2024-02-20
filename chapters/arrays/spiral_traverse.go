package main

import (
	"math"
)

const AlreadySeen = math.MinInt

func SpiralTraverse(array [][]int) []int {
	colSize := len(array)
	lineSize := len(array[0])
	totalElements := lineSize * colSize
	var ret []int

	i := 0
	matrixI, matrixJ := 0, 0
	for i < totalElements {
		// go right
		for matrixJ <= lineSize-1 && array[matrixI][matrixJ] != AlreadySeen {
			ret = append(ret, array[matrixI][matrixJ])
			array[matrixI][matrixJ] = AlreadySeen
			matrixJ++
			i++
		}
		matrixJ--
		matrixI++

		// go down
		for matrixI <= colSize-1 && array[matrixI][matrixJ] != AlreadySeen {
			ret = append(ret, array[matrixI][matrixJ])
			array[matrixI][matrixJ] = AlreadySeen
			matrixI++
			i++
		}
		matrixI--
		matrixJ--

		// go left
		for matrixJ >= 0 && array[matrixI][matrixJ] != AlreadySeen {
			ret = append(ret, array[matrixI][matrixJ])
			array[matrixI][matrixJ] = AlreadySeen
			matrixJ--
			i++
		}
		matrixJ++
		matrixI--

		// go up
		for matrixI >= 0 && array[matrixI][matrixJ] != AlreadySeen {
			ret = append(ret, array[matrixI][matrixJ])
			array[matrixI][matrixJ] = AlreadySeen
			matrixI--
			i++
		}
		matrixI++
		matrixJ++
	}

	return ret
}

//func main() {
//	fmt.Println(SpiralTraverse([][]int{
//		{1, 2, 3, 4},
//		{12, 13, 14, 5},
//		{11, 16, 15, 6},
//		{10, 9, 8, 7},
//	}))
//}
