package main

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree.Left == nil && tree.Right == nil {
		return append(array, tree.Value)
	}

	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}

	array = append(array, tree.Value)

	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}

	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}

	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}

	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}

	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}

	return append(array, tree.Value)
}

//func main() {
//	tree := &BST{Value: 10}
//	tree.Left = &BST{Value: 5}
//	tree.Left.Left = &BST{Value: 2}
//	tree.Left.Left.Left = &BST{Value: 1}
//	tree.Left.Right = &BST{Value: 5}
//	tree.Right = &BST{Value: 15}
//	tree.Right.Right = &BST{Value: 22}
//
//	fmt.Println(tree.PostOrderTraverse([]int{}))
//}
