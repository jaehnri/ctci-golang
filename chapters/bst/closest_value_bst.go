package main

import "math"

//type BST struct {
//	Value int
//
//	Left  *BST
//	Right *BST
//}

func (tree *BST) FindClosestValue(target int) int {
	return FindClosestValueRec(*tree, target, tree.Value)
}

func FindClosestValueRec(currentNode BST, target, closestValue int) int {
	if currentNode.Value == target {
		return target
	}

	closestValue = getClosest(closestValue, currentNode.Value, target)
	if currentNode.Left == nil && currentNode.Right == nil {
		return closestValue
	}

	if target < currentNode.Value && currentNode.Left != nil {
		return FindClosestValueRec(*currentNode.Left, target, closestValue)
	}

	if target > currentNode.Value && currentNode.Right != nil {
		return FindClosestValueRec(*currentNode.Right, target, closestValue)
	}

	return closestValue
}

func getClosest(a, b, target int) int {
	if math.Abs(float64(target-a)) < math.Abs(float64(target-b)) {
		return a
	}

	return b
}
