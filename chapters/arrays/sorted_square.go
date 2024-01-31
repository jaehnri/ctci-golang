package main

func SortedSquaredArray(array []int) []int {
	smallestIndex := FindSmallestIndex(array)

	leftIndex := smallestIndex - 1
	rightIndex := smallestIndex + 1
	sortedSquareArray := []int{array[smallestIndex] * array[smallestIndex]}

	for len(sortedSquareArray) != len(array) {
		if leftIndex == -1 && rightIndex != len(array) {
			sortedSquareArray = append(sortedSquareArray, array[rightIndex]*array[rightIndex])
			rightIndex++
			continue
		}

		if rightIndex == len(array) && leftIndex != -1 {
			sortedSquareArray = append(sortedSquareArray, array[leftIndex]*array[leftIndex])
			leftIndex--
			continue
		}

		if array[leftIndex]*array[leftIndex] < array[rightIndex]*array[rightIndex] {
			sortedSquareArray = append(sortedSquareArray, array[leftIndex]*array[leftIndex])
			leftIndex--
			continue
		}

		sortedSquareArray = append(sortedSquareArray, array[rightIndex]*array[rightIndex])
		rightIndex++
	}

	return sortedSquareArray
}

func FindSmallestIndex(array []int) int {
	smallestIndex := 0
	for smallestIndex < len(array)-1 {
		current := array[smallestIndex]
		next := array[smallestIndex+1]

		if next*next > current*current {
			return smallestIndex
		}

		smallestIndex++
	}

	return smallestIndex

}

//func main() {
//	fmt.Println(SortedSquaredArray([]int{1, 2, 3}))
//	fmt.Println(SortedSquaredArray([]int{-3, -2, -1}))
//	fmt.Println(SortedSquaredArray([]int{-3, 1, 4}))
//}
