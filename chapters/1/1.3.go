package main

import (
	"fmt"
	"strings"
)

// --------------------------------------------------------------------------------------------------
// 1.3) URLify: Write a method to replace all spaces in a string with '%20'. You may assume
//		that the string has sufficient space at the end to hold the additional characters,
//		and that you are given the "true" length of the string
// --------------------------------------------------------------------------------------------------

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

func main() {
	fmt.Println(replaceWithStringBuilder("Mr 3ohn Smith"))
	fmt.Println(replaceWithStringBuilder("abcd"))
	fmt.Println(replaceWithStringBuilder("abc d e"))
	fmt.Println(replaceWithStringBuilder(""))
	fmt.Println(replaceWithStringBuilder(" "))

	fmt.Println(string(replaceWithByteArray([]rune("Mr 3ohn Smith"))))
	fmt.Println(string(replaceWithByteArray([]rune("abcd"))))
	fmt.Println(string(replaceWithByteArray([]rune("abc d e"))))
	fmt.Println(string(replaceWithByteArray([]rune(""))))
	fmt.Println(string(replaceWithByteArray([]rune(" "))))
}
