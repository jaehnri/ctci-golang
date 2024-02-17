package main

import "fmt"

func SymmetricalTree(tree *BinaryTree) bool {
	if tree.Left != nil && tree.Right != nil {
		return SymmetricalTreeRec(tree.Left, tree.Right)
	}

	if tree.Left == nil && tree.Right == nil {
		return true
	}

	return false
}

func SymmetricalTreeRec(left, right *BinaryTree) bool {
	if left.Left == nil && right.Right != nil || left.Left != nil && right.Right == nil {
		return false
	}
	if right.Left == nil && left.Right != nil || right.Left != nil && left.Right == nil {
		return false
	}

	leftEqual := true
	rightEqual := true
	if left.Left != nil {
		leftEqual = SymmetricalTreeRec(left.Left, right.Right)
	}

	if left.Value != right.Value {
		return false
	}

	if left.Right != nil {
		rightEqual = SymmetricalTreeRec(left.Right, right.Left)
	}

	return leftEqual && rightEqual
}

func main() {
	tree := &BinaryTree{Value: 1}
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: 2}
	tree.Left.Left = &BinaryTree{Value: 3}
	tree.Left.Right = &BinaryTree{Value: 4}
	tree.Right.Left = &BinaryTree{Value: 4}
	tree.Right.Right = &BinaryTree{Value: 3}
	fmt.Println(SymmetricalTree(tree))
}
