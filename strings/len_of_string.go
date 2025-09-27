package main

import (
	"strings"
)

func LenOfString(s string) int {
	var arry []string = strings.Split(s, " ")
	var last_word string = arry[len(arry)-1]

	return len(last_word)
}
