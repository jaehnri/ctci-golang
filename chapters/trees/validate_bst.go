package main

import (
	"fmt"
)

type BSTInfo struct {
	Valid bool
	Max   int
	Min   int
}

func (tree *BST) ValidateBst() bool {
	return tree.ValidateBstMaxMin().Valid
}

func (tree *BST) ValidateBstMaxMin() BSTInfo {
	if tree.Left == nil && tree.Right == nil {
		return BSTInfo{Min: tree.Value, Max: tree.Value, Valid: true}
	}

	var treeMax int
	var treeMin int

	if tree.Left != nil {
		leftBst := tree.Left.ValidateBstMaxMin()
		if !leftBst.Valid || !(tree.Value > leftBst.Max) {
			return BSTInfo{Valid: false}
		}
		treeMin = leftBst.Min
	} else {
		treeMin = tree.Value
	}

	if tree.Right != nil {
		rightBst := tree.Right.ValidateBstMaxMin()
		if !rightBst.Valid || !(tree.Value <= rightBst.Min) {
			return BSTInfo{Valid: false}
		}
		treeMax = rightBst.Max
	} else {
		treeMax = tree.Value
	}

	return BSTInfo{Valid: true, Min: treeMin, Max: treeMax}
}

func main() {
	tree := &BST{Value: 10}
	tree.Left = &BST{Value: 5}
	tree.Left.Left = &BST{Value: 2}
	tree.Left.Left.Left = &BST{Value: 1}
	tree.Left.Right = &BST{Value: 5}
	tree.Left.Right.Right = &BST{Value: 11}
	tree.Right = &BST{Value: 15}
	tree.Right.Right = &BST{Value: 22}

	fmt.Println(tree.ValidateBstMaxMin())
}
