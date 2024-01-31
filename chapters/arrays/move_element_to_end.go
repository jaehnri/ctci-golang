package main

func MoveElementToEnd(array []int, toMove int) []int {
	targetIdx := 0
	notTargetIdx := len(array) - 1

	for targetIdx < notTargetIdx {
		if array[targetIdx] != toMove {
			targetIdx++
			continue
		}

		if array[notTargetIdx] == toMove {
			notTargetIdx--
			continue
		}

		aux := array[targetIdx]
		array[targetIdx] = array[notTargetIdx]
		array[notTargetIdx] = aux
	}

	return nil
}
