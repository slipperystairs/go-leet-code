// Given an array of strings "words", return the first palindromic
// string in the array. If there is no such string, return an empty string.
// A string is palindromic if it reads the same forward and backwards.
package main

func isValid(word string) bool {
	// Checking for plain ASCII works, but runes are cooler
	// right := len(word) - 1
	var runes []rune = []rune(word)
	var right int = len(runes) - 1
	var left int = 0

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func IsPalindrome(words []string) string {
	for _, word := range words {
		if isValid(word) {
			return word
		}
	}
	return ""
}
