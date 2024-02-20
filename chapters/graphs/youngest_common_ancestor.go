package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	onePointer := descendantOne
	twoPointer := descendantTwo

	oneDescendantHeight := 0
	for onePointer != top {
		oneDescendantHeight++
		onePointer = onePointer.Ancestor
	}

	twoDescendantHeight := 0
	for twoPointer != top {
		twoDescendantHeight++
		twoPointer = twoPointer.Ancestor
	}

	onePointer = descendantOne
	twoPointer = descendantTwo
	for twoDescendantHeight > oneDescendantHeight {
		twoPointer = twoPointer.Ancestor
		twoDescendantHeight--
	}

	for oneDescendantHeight > twoDescendantHeight {
		onePointer = onePointer.Ancestor
		oneDescendantHeight--
	}

	for onePointer.Name != twoPointer.Name {
		onePointer = onePointer.Ancestor
		twoPointer = twoPointer.Ancestor
	}

	return onePointer
}
