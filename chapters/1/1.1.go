package main

import (
	"fmt"
	"strings"
)

// --------------------------------------------------------------------------------------------------
// 1.1) Is Unique: Implement an algorithm to determine if a string has all unique characters.
//      What if you  cannot use additional data structures?
//
// R.: If you can't use additional data structures, we could consider using two options:
//     1) If we are able to modify the string, we could use a heap sort. Heap sort is possible
//		  because it requires no additional space. Then, we could check for neighbor characters
//		  that are equal. Heap sort runs in O(n * log n) and the iteration in O(n).
//		  Thus, it would run in an O(n * log N) time complexity.
//     2) If we aren't able to modify the string and create extra data structures, our only option is
//        to compare letter by letter. First letter would compare itself with the other n - 1,
//        the second would compare itself with another n - 2, etc.
//        This would result in O(n - 1 + n - 2 + ... + 1),
//        i.e, O( (1 + n - 1) * (n - 1)/2 ) = O(n * (n - 1)/2 ) = O((nˆ2 - n)/2 ) = O(nˆ2)
// --------------------------------------------------------------------------------------------------

// isUnique runs in O(n). Hashmap insertions and searches are amortized O(1) and we perform 'n' of these
// searches in the worst case, 'n' being the length of the string s.
func isUnique(s string) bool {

	// Assuming ASCII alphabet (128 unique characters), a string with size greater than
	// 128 must have repeated characters.
	if len(s) > 128 {
		return false
	}

	m := make(map[rune]int)
	for _, char := range s {

		// m[char] == 1 means that 'char' already appeared in string.
		if m[char] == 1 {
			return false
		} else {
			m[char] = 1
		}
	}

	return true
}

func stringWithAllAsciiCharacters() string {
	var sb strings.Builder
	i := 0
	for i < 128 {
		sb.WriteString(string(i))
		i++
	}

	return sb.String()
}

func main() {
	fmt.Println(isUnique("abc"))
	fmt.Println(isUnique("abcc"))
	fmt.Println(isUnique(stringWithAllAsciiCharacters()))
	fmt.Println(isUnique(stringWithAllAsciiCharacters() + "a"))
}
