package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head   *Node
	Length int
}

func (l *List) Insert(n *Node) {

	currentNode := l.Head

	if currentNode == nil {
		l.Head = n
		l.Length++
		return
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = n
	l.Length++
}

func (l *List) Remove(data int) {
	currentNode := l.Head

	if currentNode.Data == data {
		l.Head = currentNode.Next
		return
	}

	prev := l.Head
	currentNode = currentNode.Next

	for currentNode != nil {
		if currentNode.Data == data {
			prev.Next = currentNode.Next
			l.Length--
		}

		prev = currentNode
		currentNode = currentNode.Next
	}
}

func (l *List) Print() {
	currentNode := l.Head

	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
}

func main() {
	var list List

	list.Insert(&Node{Data: 1, Next: nil})
	list.Insert(&Node{Data: 2, Next: nil})
	list.Insert(&Node{Data: 3, Next: nil})
	list.Insert(&Node{Data: 4, Next: nil})
	list.Insert(&Node{Data: 4, Next: nil})
	list.Print()

	fmt.Println()
	list.Remove(4)
	list.Print()
}

// 1 2 3 4
// 1 2 4
// tem que fazer o prev apontar pro next
