package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	numAsStr := strconv.Itoa(x)
	middleString := ""

	if len(numAsStr) == 1 {
		return true
	}

	for index, letter := range numAsStr {
		switch index {
		case 0:
			continue
		case len(numAsStr) - 1:
			continue
		default:
			middleString = string(letter) + middleString
		}
	}

	newString := numAsStr[len(numAsStr)-1:] + middleString + numAsStr[0:1]
	return newString == numAsStr
}

func main() {
	fmt.Println(isPalindrome(121))
}
