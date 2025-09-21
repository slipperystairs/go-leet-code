package main

func RomanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := 0

	for i, val := range s {
		// If current value is less than next value, subtract instead of add
		if i+1 < len(s) && romanMap[val] < romanMap[rune(s[i+1])] {
			num -= romanMap[val]
		} else {
			num += romanMap[val]
		}
	}

	return num
}
