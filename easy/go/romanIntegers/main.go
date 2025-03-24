package main

import "fmt"

func romanToInts(s string) int {
	romanIntConversion := 0
	romanIntegers := make(map[string]int)
	romanIntegers["I"] = 1
	romanIntegers["V"] = 5
	romanIntegers["X"] = 10
	romanIntegers["L"] = 50
	romanIntegers["C"] = 100
	romanIntegers["D"] = 500
	romanIntegers["M"] = 1000
	skipNext := false
	lengthOfString := len(s)

	for index, letter := range s {
		if skipNext {
			skipNext = false
			continue
		}
		if string(letter) == "I" && index+1 < lengthOfString && string(s[index+1]) == "V" {
			romanIntConversion += 4
			skipNext = true
		} else if string(letter) == "I" && index+1 < lengthOfString && string(s[index+1]) == "X" {
			romanIntConversion += 9
			skipNext = true
		} else if string(letter) == "X" && index+1 < lengthOfString && string(s[index+1]) == "L" {
			romanIntConversion += 40
			skipNext = true
		} else if string(letter) == "X" && index+1 < lengthOfString && string(s[index+1]) == "C" {
			romanIntConversion += 90
			skipNext = true
		} else if string(letter) == "C" && index+1 < lengthOfString && string(s[index+1]) == "D" {
			romanIntConversion += 400
			skipNext = true
		} else if string(letter) == "C" && index+1 < lengthOfString && string(s[index+1]) == "M" {
			romanIntConversion += 900
			skipNext = true
		} else {
			romanIntConversion += romanIntegers[string(letter)]
			skipNext = false
		}
	}

	return romanIntConversion
}

func main() {
	fmt.Println(romanToInts("III"))
}
