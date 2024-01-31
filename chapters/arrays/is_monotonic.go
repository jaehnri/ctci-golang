package main

func IsMonotonic(array []int) bool {
	nonDecreasing := true
	nonIncreasing := true

	i := 0
	for i < len(array)-1 {
		if array[i+1] < array[i] {
			nonDecreasing = false
		}
		if array[i+1] > array[i] {
			nonIncreasing = false
		}
		i++
	}

	return nonIncreasing || nonDecreasing

}

//func main() {
//	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))
//}
