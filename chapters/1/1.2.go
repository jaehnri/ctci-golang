package main

import "fmt"

// --------------------------------------------------------------------------------------------------
// 1.2) Check Permutation: Given two strings, write a method to decide if one is a permutation of the
//
//	other.
//
// --------------------------------------------------------------------------------------------------

// createKeysAndHashmapOfString runs in O(n), n being the size of s. Queries and insertions in a hashmap run in O(1).
func createKeysAndHashmapOfString(s string) map[rune]int {
	m := make(map[rune]int)

	for _, char := range s {
		value, ok := m[char]
		if ok {
			m[char] = value + 1
		} else {
			m[char] = 1
		}
	}

	return m
}

// checkPermutation runs in O(n). It calls createKeysAndHashmapOfString twice (2 * O(n)) and compares
// the hashmaps key by key.
// If we consider that the size of these strings is n, we do 2 * n queries, comparing then.s
func checkPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := createKeysAndHashmapOfString(s1)
	m2 := createKeysAndHashmapOfString(s2)

	for key := range m1 {
		if m2[key] != m1[key] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkPermutation("abc", "cab"))
	fmt.Println(checkPermutation("abc", "cabd"))
	fmt.Println(checkPermutation("abc", "cad"))
}
