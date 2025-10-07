package main

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	t.Run("testing merging two string alternately", func(t *testing.T) {
		var test = []struct {
			word1  string
			word2  string
			result string
		}{
			{"abc", "ch", "acbhc"},
			{"ab", "pqrs", "apbqrs"},
			{"abcd", "pq", "apbqcd"},
		}

		for _, test := range test {
			var result string = MergeAlternately(test.word1, test.word2)
			if result != test.result {
				t.Errorf("Got %s but wanted %s\n", result, test.result)
			} else {
				fmt.Printf("Pass %s\n", result)
			}
		}
	})
}
