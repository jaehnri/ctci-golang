package main

func HasSingleCycle(array []int) bool {
	visited := make([]bool, len(array))

	count := 0
	current := 0

	for count < len(array) {
		step := array[current]
		current = updateIndex(current, step, len(array))
		if visited[current] == true {
			return false
		}
		visited[current] = true
		count++
	}

	for _, value := range visited {
		if value == false {
			return false
		}
	}

	return true
}

// current is the current index of the array and is always positive
// offset is the step
// bound is meant to be the length of the array
// 0 + 1, 3 => 1
// 0 + 3, 3 => 0
// 0 + 4, 3 => 1
// 0 - 1, 3 => 2
// 0 - 2, 3 => 1
// 0 + 10, 3 => 0 + 7 => 0 + 4 => 0 + 1
// 0 - 10 => 0 - 7 => 0 - 4 => 0 - 1 => 2
func updateIndex(current, offset, bound int) int {
	for offset >= bound {
		offset -= bound
	}

	for offset <= -bound {
		offset += bound
	}

	if current+offset < 0 {
		return bound + (current + offset)
	}
	if current+offset < bound {
		return current + offset
	}

	return current + offset - bound
}
