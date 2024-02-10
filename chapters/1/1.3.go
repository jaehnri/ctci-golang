package main

import (
	"strings"
)

// --------------------------------------------------------------------------------------------------
// 1.3) URLify: Write a method to replace all spaces in a string with '%20'. You may assume
//		that the string has sufficient space at the end to hold the additional characters,
//		and that you are given the "true" length of the string
// --------------------------------------------------------------------------------------------------

// replaceWithStringBuilder runs in O(n) (n is the size of the new string) as WriteRune and WriteString
// only append a letter to an array.
// This operation is O(1), becoming O(n) only when the array size needs to be increased, i.e, a bigger
// buffer is created and the contents of the old full buffer needs to be copied. This rarely happen, so
// the amortized cost is still O(1). sb.String() is also O(n), but is performed only once.
// Space complexity is also O(n) as strings.Builder creates an additional buffer.
func replaceWithStringBuilder(s string) string {
	var sb strings.Builder

	for _, char := range s {
		if char == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

func calculateNewSize(s []rune) int {
	spaceCount := 0

	for _, char := range s {
		if char == ' ' {
			spaceCount++
		}
	}

	return len(s) + 2*spaceCount
}

// replaceWithByteArray runs in O(n) too. First, a scan in the string s is done to see how many spaces
// there are and check the size of the new string.
// Then, the creation of the new string is done in reverse order. Space complexity is O(n) since a
// bigger buffer is created, n being the size of the new string.
func replaceWithByteArray(s []rune) []rune {
	originalSize := len(s)
	newSize := calculateNewSize(s)

	if originalSize == newSize {
		return s
	}

	newString := make([]rune, newSize)

	originalStringIndex := originalSize - 1
	for i := newSize - 1; i >= 0; i-- {
		if s[originalStringIndex] == ' ' {
			newString[i] = '0'
			newString[i-1] = '2'
			newString[i-2] = '%'
			i = i - 2
		} else {
			newString[i] = s[originalStringIndex]
		}

		originalStringIndex--
	}

	return newString
}

// func main() {
// 	fmt.Println(replaceWithStringBuilder("Mr 3ohn Smith"))
// 	fmt.Println(replaceWithStringBuilder("abcd"))
// 	fmt.Println(replaceWithStringBuilder("abc d e"))
// 	fmt.Println(replaceWithStringBuilder(""))
// 	fmt.Println(replaceWithStringBuilder(" "))

// 	fmt.Println(string(replaceWithByteArray([]rune("Mr 3ohn Smith"))))
// 	fmt.Println(string(replaceWithByteArray([]rune("abcd"))))
// 	fmt.Println(string(replaceWithByteArray([]rune("abc d e"))))
// 	fmt.Println(string(replaceWithByteArray([]rune(""))))
// 	fmt.Println(string(replaceWithByteArray([]rune(" "))))
// }
