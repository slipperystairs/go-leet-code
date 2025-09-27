package main

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Run("palindrome test", func(t *testing.T) {
		want := "aba"
		fmt.Printf("Want: %s\n", want)
		strings := []string{"aba", "abc"}
		got := IsPalindrome(strings)
		fmt.Printf("Got: %s\n", got)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
