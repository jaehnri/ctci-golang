package main

import "errors"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if tree == nil {
		tree.Value = value
		return tree
	}

	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			return tree.Left.Insert(value)
		}
	}

	if value >= tree.Value {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			return tree.Right.Insert(value)
		}
	}

	return tree
}

func (tree *BST) Contains(value int) bool {
	if tree == nil {
		return false
	}

	if tree.Value == value {
		return true
	}

	if value < tree.Value {
		return tree.Left.Contains(value)
	}

	return tree.Right.Contains(value)
}

func (tree *BST) Remove(value int) *BST {
	// Don't remove if it is a single node tree
	if value == tree.Value && tree.Left == nil && tree.Right == nil {
		return tree
	}

	var parent *BST
	nodeToRemove := tree

	found := false
	for !found {
		if nodeToRemove == nil {
			return tree
		}

		if value == nodeToRemove.Value {
			found = true
			continue
		}

		if value < nodeToRemove.Value {
			parent = nodeToRemove
			nodeToRemove = nodeToRemove.Left
			continue
		}

		if value > nodeToRemove.Value {
			parent = nodeToRemove
			nodeToRemove = nodeToRemove.Right
		}
	}

	// 1) Node is leaf, we can just remove it
	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return tree
	}

	// 2) Node has only one child and can be replaced
	if nodeToRemove.Left == nil && nodeToRemove.Right != nil {
		if parent == nil {
			tree.Value = tree.Right.Value
			tree.Right = tree.Right.Right
			return tree
		}
		parent.Right = nodeToRemove.Right
		return tree.Right
	}
	if nodeToRemove.Right == nil && nodeToRemove.Left != nil {
		parent.Left = nodeToRemove.Left
		return tree
	}

	// 3) Node has two children and has to be replaced by min of right tree
	if nodeToRemove.Left != nil && nodeToRemove.Right != nil {
		var successorParent *BST
		successor := nodeToRemove.Right
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		nodeToRemove.Value = successor.Value
		if successorParent != nil {
			successorParent.Left = nil
		} else {
			// there's no parent, removed node is root
			tree.Right = nil
		}
		return tree
	}

	return tree
}

func (tree *BST) Find(value int) (*BST, error) {
	if tree == nil {
		return nil, errors.New("value not in tree")
	}

	if value == tree.Value {
		return tree, nil
	}

	if value < tree.Value {
		return tree.Left.Find(value)
	}

	return tree.Right.Find(value)
}

//func main() {
//	tree := &BST{Value: 1}
//	tree.Insert(2)
//	tree.Insert(3)
//	tree.Insert(4)
//	tree.Remove(1)
//}
