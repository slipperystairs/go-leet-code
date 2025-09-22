package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("binary search test", func(t *testing.T) {
		numbers := []int{-1, 0, 3, 5, 9, 12}
		target := 9
		want := 4
		got := Search(numbers, target)
		fmt.Printf("got %d\n", got)

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
