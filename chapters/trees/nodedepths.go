package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return root.NodeDepthsRec(0)
}

func (bt *BinaryTree) NodeDepthsRec(currentDepth int) int {
	if bt.Left == nil && bt.Right == nil {
		return currentDepth
	}

	var leftSum int
	var rightSum int

	if bt.Left != nil {
		leftSum = bt.Left.NodeDepthsRec(currentDepth + 1)
	}

	if bt.Right != nil {
		rightSum = bt.Right.NodeDepthsRec(currentDepth + 1)
	}

	return currentDepth + leftSum + rightSum
}

//func main() {
//	root := &BinaryTree{Value: 1}
//	root.Left = &BinaryTree{Value: 2}
//	root.Left.Left = &BinaryTree{Value: 4}
//	root.Left.Left.Left = &BinaryTree{Value: 8}
//	root.Left.Left.Right = &BinaryTree{Value: 9}
//	root.Left.Right = &BinaryTree{Value: 5}
//	root.Right = &BinaryTree{Value: 3}
//	root.Right.Left = &BinaryTree{Value: 6}
//	root.Right.Right = &BinaryTree{Value: 7}
//	fmt.Println(NodeDepths(root))
//}
