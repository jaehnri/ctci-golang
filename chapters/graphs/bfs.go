package main

import (
	"errors"
)

type Queue struct {
	fifo []*Node
}

func (q *Queue) enqueue(a *Node) {
	q.fifo = append(q.fifo, a)
}

func (q *Queue) dequeue() (*Node, error) {
	if len(q.fifo) == 0 {
		return nil, errors.New("queue is empty")
	}

	defer func() { q.fifo = q.fifo[1:] }()

	return q.fifo[0], nil
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	q := Queue{}
	currentNode := n
	for currentNode != nil {
		array = append(array, currentNode.Name)
		for _, child := range currentNode.Children {
			q.enqueue(child)
		}

		currentNode, _ = q.dequeue()
	}

	return array
}
