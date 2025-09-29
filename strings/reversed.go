package main

func Solution(word string) string {
	var reversed string = ""

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	return reversed
}
