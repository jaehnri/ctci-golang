package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	oneAncestors := make(map[string]bool)
	twoAncestors := make(map[string]bool)

	onePointer := descendantOne
	twoPointer := descendantTwo

	for onePointer != nil || twoPointer != nil {
		oneAncestors[onePointer.Name] = true
		twoAncestors[twoPointer.Name] = true

		if oneAncestors[twoPointer.Name] == true {
			return twoPointer
		}

		if twoAncestors[onePointer.Name] == true {
			return onePointer
		}

		if onePointer.Ancestor != nil {
			onePointer = onePointer.Ancestor
		}

		if twoPointer.Ancestor != nil {
			twoPointer = twoPointer.Ancestor
		}
	}

	return top
}
