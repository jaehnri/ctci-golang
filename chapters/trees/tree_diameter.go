package main

func BinaryTreeDiameter(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	if tree.Left == nil && tree.Right == nil {
		return 0
	}

	_, diameter := tree.BinaryTreeDiameterRec()
	return diameter
}

// BinaryTreeDiameterRec returns two ints. First is the size from that node to the leaf
// and the second is the max diameter so far.
func (bt *BinaryTree) BinaryTreeDiameterRec() (length int, max int) {
	if bt.Left == nil && bt.Right == nil {
		return 0, 0
	}

	var leftLength, leftMax, rightLength, rightMax int
	if bt.Left != nil {
		leftLength, leftMax = bt.Left.BinaryTreeDiameterRec()
		leftLength++
	}

	if bt.Right != nil {
		rightLength, rightMax = bt.Right.BinaryTreeDiameterRec()
		rightLength++
	}

	length = Max(rightLength, leftLength)
	max = Max(Max(Max(leftLength+rightLength, leftMax), rightMax), length)

	return length, max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//func main() {
//	//root := &BinaryTree{Value: 1}
//	//root.Left = &BinaryTree{Value: 3}
//	//root.Left.Left = &BinaryTree{Value: 7}
//	//root.Left.Left.Left = &BinaryTree{Value: 8}
//	//root.Left.Left.Left.Right = &BinaryTree{Value: 1}
//	//root.Left.Left.Left.Right.Right = &BinaryTree{Value: 2}
//	//root.Left.Left.Left.Left = &BinaryTree{Value: 9}
//	//root.Left.Right = &BinaryTree{Value: 4}
//	//root.Left.Right.Right = &BinaryTree{Value: 5}
//	//root.Left.Right.Right.Right = &BinaryTree{Value: 6}
//	//root.Right = &BinaryTree{Value: 2}
//	//fmt.Println(BinaryTreeDiameter(root))
//
//	//root := &BinaryTree{Value: 4}
//	//root.Left = &BinaryTree{Value: 2}
//	//fmt.Println(BinaryTreeDiameter(root))
//}
