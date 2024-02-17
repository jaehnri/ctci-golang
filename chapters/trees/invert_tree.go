package main

func (bt *BinaryTree) InvertBinaryTree() {
	if bt == nil {
		return
	}

	if bt.Left == nil && bt.Right == nil {
		return
	}

	if bt.Left != nil {
		bt.Left.InvertBinaryTree()
	}

	if bt.Right != nil {
		bt.Right.InvertBinaryTree()
	}

	rightTree := bt.Right
	leftTree := bt.Left

	bt.Right = leftTree
	bt.Left = rightTree
}

//func main() {
//	tree := &BinaryTree{Value: 1}
//	tree.Left = &BinaryTree{Value: 2}
//	tree.Right = &BinaryTree{Value: 3}
//	tree.Left.Left = &BinaryTree{Value: 4}
//	tree.Left.Right = &BinaryTree{Value: 5}
//	tree.Left.Left.Left = &BinaryTree{Value: 8}
//	tree.Left.Left.Right = &BinaryTree{Value: 9}
//	tree.Right.Left = &BinaryTree{Value: 6}
//	tree.Right.Right = &BinaryTree{Value: 7}
//	tree.InvertBinaryTree()
//}
