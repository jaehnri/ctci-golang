package main

type BinaryTreeWithParent struct {
	Value int

	Left   *BinaryTreeWithParent
	Right  *BinaryTreeWithParent
	Parent *BinaryTreeWithParent
}

func FindSuccessor(tree *BinaryTreeWithParent, node *BinaryTreeWithParent) *BinaryTreeWithParent {
	if tree.Left != nil {
		successor := FindSuccessor(tree.Left, node)
		if successor != nil {
			return successor
		}
	}

	if node == tree {
		if tree.Right != nil {
			successor := tree.Right
			for successor.Left != nil {
				successor = successor.Left
			}
			return successor
		}

		if tree.Parent != nil && tree.Parent.Left == tree {
			return tree.Parent
		}

		if tree.Parent != nil && tree.Parent.Parent != nil && tree.Parent.Right == tree {
			return tree.Parent.Parent
		}
	}

	if tree.Right != nil {
		successor := FindSuccessor(tree.Right, node)
		if successor != nil {
			return successor
		}
	}

	return nil
}

//func main() {
//	root := &BinaryTreeWithParent{Value: 1}
//	root.Left = &BinaryTreeWithParent{Value: 2}
//	root.Left.Parent = root
//	root.Right = &BinaryTreeWithParent{Value: 3}
//	root.Right.Parent = root
//	root.Left.Left = &BinaryTreeWithParent{Value: 4}
//	root.Left.Left.Parent = root.Left
//	root.Left.Right = &BinaryTreeWithParent{Value: 5}
//	root.Left.Right.Parent = root.Left
//	root.Left.Left.Left = &BinaryTreeWithParent{Value: 6}
//	root.Left.Left.Left.Parent = root.Left.Left
//	fmt.Println(FindSuccessor(root, root.Left.Right))
//}
