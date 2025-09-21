package main

import "testing"

func TestRomanToInt(t *testing.T) {
	t.Run("roman numeral 4", func(t *testing.T) {
		got := RomanToInt("IV")
		want := 4

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
