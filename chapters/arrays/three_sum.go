package main

import (
	"fmt"
	"slices"
)

func ThreeNumberSum(array []int, target int) [][]int {
	slices.Sort(array)
	possibleSums := [][]int{}

	for index, number := range array {
		l := index + 1
		r := len(array) - 1

		for l < r {
			leftValue := array[l]
			rightValue := array[r]
			if number+array[l]+rightValue == target {
				sum := []int{number, leftValue, rightValue}
				possibleSums = append(possibleSums, sum)
				l++
				r--
				continue
			}

			if number+leftValue+rightValue > target {
				r--
				continue
			}

			if number+leftValue+rightValue < target {
				l++
				continue
			}
		}
	}

	return possibleSums
}

func main() {
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
}
