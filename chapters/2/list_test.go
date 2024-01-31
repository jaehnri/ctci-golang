package main

import (
	"testing"
)

func TestRemoveOnlyOneOccurrence(t *testing.T) {
	l := createList([]int{1, 2, 3, 3})

	l.Remove(3)
	if l.Length != 3 {
		t.Fatalf("Not only one occurence of an item was removded.")
	}
}

func createList(data []int) List {
	var l List
	for _, value := range data {
		l.Insert(&Node{Data: value, Next: nil})
	}

	return l
}
