package main

import "fmt"

// --------------------------------------------------------------------------------------------------
// 1.4) Palindrome Permutation: Given a string, write a function to check if it is a permutation of
//      a palindrome. A palindrome is a word or phrase that is the same forwards and backwards.
//      A permutation is a rearrangement of letters. The palindrome does not need to be limited to
//      just dictionary words.
// --------------------------------------------------------------------------------------------------

func isOdd(a int) bool {
	return a%2 != 0
}

func isPermutationOfPalindrome(s string) bool {
	m := make(map[rune]int)

	for _, char := range s {
		val, ok := m[char]

		if ok {
			m[char] = val + 1
		} else {
			m[char] = 1
		}
	}

	odds := 0
	for key := range m {
		if isOdd(m[key]) {
			odds++
		}
	}

	return isOdd(odds) || odds == 0
}

func main() {
	fmt.Println(isPermutationOfPalindrome("a"))
	fmt.Println(isPermutationOfPalindrome("aa"))
	fmt.Println(isPermutationOfPalindrome("aab"))
	fmt.Println(isPermutationOfPalindrome("aaaab"))
	fmt.Println(isPermutationOfPalindrome("aabbb"))
	fmt.Println(isPermutationOfPalindrome("ab"))
	fmt.Println(isPermutationOfPalindrome("aaab"))
	fmt.Println(isPermutationOfPalindrome("tactcoapapa"))
}
