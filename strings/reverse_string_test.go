package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReversedString(t *testing.T) {
	t.Run("test reversed byte of characters", func(t *testing.T) {
		var tests = []struct {
			input    []byte
			expected []byte
		}{
			{[]byte("hello"), []byte("olleh")},
			{[]byte("racecar"), []byte("racecar")},
			{[]byte("abcd"), []byte("dcba")},
		}

		for _, test := range tests {
			// copy to avoid copy mutation side effects
			var result []byte = ReversedString(append([]byte(nil), test.input...))
			fmt.Printf("result: %s\n", string(result))

			if !bytes.Equal(result, test.expected) {
				t.Errorf("Got %q but we wanted %q ", result, test.expected)
			}
		}
	})

}
