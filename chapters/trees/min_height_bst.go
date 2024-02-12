package main

func MinHeightBST(array []int) *BST {
	return MinHeightRec(array, nil)
}

func MinHeightRec(array []int, bst *BST) *BST {
	if len(array) == 0 {
		return bst
	}

	if len(array) == 1 {
		bst.Insert(array[0])
		return bst
	}

	toInsert := len(array) / 2
	leftSubArray := array[:toInsert]
	rightSubArray := array[toInsert+1:]

	bst.Insert(array[toInsert])
	bst = MinHeightRec(leftSubArray, bst)
	bst = MinHeightRec(rightSubArray, bst)

	return bst
}

//func main() {
//	bst := MinHeightBST([]int{1, 2, 5, 7, 10, 13, 14, 15, 22})
//	fmt.Println(bst)
//}
