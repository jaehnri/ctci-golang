package main

func FindKthLargestValueInBst(tree *BST, k int) int {
	var array []int
	array = ReverseOrderTraversal(tree, array)
	return array[k-1]
}

func ReverseOrderTraversal(tree *BST, array []int) []int {
	if tree.Right != nil {
		array = ReverseOrderTraversal(tree.Right, array)
	}

	array = append(array, tree.Value)

	if tree.Left != nil {
		array = ReverseOrderTraversal(tree.Left, array)
	}

	return array
}

func FindKthLargestValueInBst0Space(tree *BST, k int) int {
	kthLargest, _ := ReverseOrderTraversalWithoutArray(tree, k)
	return kthLargest
}

func ReverseOrderTraversalWithoutArray(tree *BST, k int) (kthLargest int, current int) {
	if tree.Right != nil {
		kthLargest, k = ReverseOrderTraversalWithoutArray(tree.Right, k)
	}

	if k == 0 {
		return kthLargest, k
	}

	if k == 1 {
		return tree.Value, 0
	}
	k--

	if tree.Left != nil {
		kthLargest, k = ReverseOrderTraversalWithoutArray(tree.Left, k)
	}

	return kthLargest, k
}

//func main() {
//	bst := &BST{Value: 15}
//	bst.Insert(5)
//	bst.Insert(20)
//	bst.Insert(2)
//	bst.Insert(5)
//	bst.Insert(17)
//	bst.Insert(22)
//	bst.Insert(1)
//	bst.Insert(3)
//	//fmt.Println(FindKthLargestValueInBst(bst, 1))
//	fmt.Println(FindKthLargestValueInBst0Space(bst, 2))
//}
