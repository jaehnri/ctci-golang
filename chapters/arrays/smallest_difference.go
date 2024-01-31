package main

import (
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	smallestDiff, _ := AbsoluteDiff(array1[0], array2[0])
	smallestPair := []int{array1[0], array2[0]}

	l, r := 0, 0
	for l < len(array1) && r < len(array2) {
		diff, bigger := AbsoluteDiff(array1[l], array2[r])
		if diff < smallestDiff {
			smallestDiff = diff
			smallestPair = []int{array1[l], array2[r]}
		}

		if bigger == 0 {
			r++
			continue
		}

		if bigger == 1 {
			l++
			continue
		}

		// bigger == 2 => difference is 0
		return smallestPair
	}

	return smallestPair
}

// AbsoluteDiff returns the absolute difference between 2 numbers and which of the two is bigger:
// 0 means a is bigger, 1 means b is bigger, and 2 means they are equal.
//
//	3, 4 => 1, 1
//	3, 3 => 0, 0
//
// -1, 4 => 5, 1
func AbsoluteDiff(a, b int) (int, int) {
	if a < b {
		return b - a, 1
	}

	if b < a {
		return a - b, 0
	}

	return a - b, 2
}

//func main() {
//	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
//}
