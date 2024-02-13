package main

import (
	"fmt"
	"math"
)

type Interval struct {
	Min int
	Max int
}

type Intervals struct {
	Left  Interval
	Right Interval
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	tree := &BST{Value: preOrderTraversalValues[0]}
	tree.ReconstructBstRec(
		preOrderTraversalValues[1:],
		0,
		Intervals{
			Left:  Interval{Min: math.MinInt, Max: tree.Value},
			Right: Interval{Min: tree.Value, Max: math.MaxInt},
		},
	)
	return tree
}

func (tree *BST) ReconstructBstRec(array []int, index int, intervals Intervals) ([]int, int) {
	if index == len(array) {
		return array, index
	}

	if array[index] < intervals.Left.Max && array[index] >= intervals.Left.Min {
		v := array[index]
		tree.Left = &BST{Value: v}
		index++
		array, index = tree.Left.ReconstructBstRec(
			array,
			index,
			Intervals{
				Left:  Interval{Min: intervals.Left.Min, Max: v},
				Right: Interval{Min: v, Max: tree.Value},
			},
		)
	}

	if index == len(array) {
		return array, index
	}

	if array[index] >= intervals.Right.Min && array[index] < intervals.Right.Max {
		v := array[index]
		tree.Right = &BST{Value: v}
		index++
		array, index = tree.Right.ReconstructBstRec(
			array,
			index,
			Intervals{
				Left:  Interval{Min: intervals.Left.Min, Max: v},
				Right: Interval{Min: v, Max: intervals.Right.Max},
			},
		)
	}

	return array, index
}

func main() {
	tree := ReconstructBst([]int{10, 4, 2, 1, 5, 17, 19, 18})
	fmt.Println(tree)
}
