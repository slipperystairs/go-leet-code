package main

import (
	"fmt"
	"testing"
)

func TestLenOfString(t *testing.T) {
	t.Run("len of string test", func(t *testing.T) {
		want := 5
		got := LenOfString("Hello World")
		fmt.Printf("got %d\n", got)

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
