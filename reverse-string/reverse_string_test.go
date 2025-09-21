package main

import "testing"

func TestSolution(t *testing.T) {
	t.Run("roman numeral 4", func(t *testing.T) {
		got := Solution("ass")
		want := "ssa"

		if got != want {
			t.Errorf("got %s want %s given", got, want)
		}
	})
}
