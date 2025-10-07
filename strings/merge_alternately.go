// Your are given two strings word1 and word2.
// Merge the strings by adding the letters in alternating order, starting with word1
// If a string is longer than the other, append the additional letters onto the end of the merged string.
// Return the merged string.

// Example 1:
// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

package main

func MergeAlternately(word1 string, word2 string) string {
	var alternating []byte

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			alternating = append(alternating, word1[i])
		}
		if i < len(word2) {
			alternating = append(alternating, word2[i])
		}
	}

	return string(alternating)
}
