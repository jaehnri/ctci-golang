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

// isPermutationOfPalindrome runs in O(n). Firstly, it creates a hashmap where the key is a char
// and the value is how many times it appears on the string.
// In order for a string to be a palindrome, it must be made of pairs of letters. At most, only
// one character can appear an odd number of times in this string, i.e., the middle character.
// After the creation of this hashmap, it iterates over the keys of the map and checks there are
// no more than one odd number
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

	return odds <= 1
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
