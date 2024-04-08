package main

import "fmt"

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	if n == 0 {
		return 1
	}

	total := 0
	numberOfWays := make([]int, n)

	for _, coin := range denoms {
		i := coin

		for i < n {
			numberOfWays[i]++

			total += numberOfWays[n-i]
			i += coin
		}

		if i == n {
			total++
		}
	}

	return total
}

func main() {
	fmt.Println(NumberOfWaysToMakeChange(6, []int{1, 5}))
}
