package main

import (
	"fmt"
	"testing"
)

func TestDup(t *testing.T) {
	t.Run("test for duplicates", func(t *testing.T) {
		numbers := []int{1, 2, 3, 1}
		want := true
		got := ContainsDuplicate(numbers)
		fmt.Printf("returned %t\n", got)

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
