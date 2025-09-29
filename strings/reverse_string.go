package main

func ReversedString(chars []byte) []byte {
	var left int = 0
	var right int = len(chars) - 1
	var temp byte

	for left < right {
		temp = chars[left]
		chars[left] = chars[right]
		chars[right] = temp
		left++
		right--
	}

	return chars
}
