package main

import (
	"testing"
)

func TestIsPalindromeCaseOne(t *testing.T) {
	num := 121
	answer := isPalindrome(num)
	expected := true
	if !answer == expected {
		t.Errorf("Answer should be %v for %v but %v was returned...", expected, num, answer)
	}
}

func TestIsPalindromeCaseTwo(t *testing.T) {
	num := -121
	answer := isPalindrome(num)
	expected := false
	if !answer == expected {
		t.Errorf("Answer should be %v for %v but %v was returned...", expected, num, answer)
	}
}

func TestIsPalindromeCaseThree(t *testing.T) {
	num := 0
	answer := isPalindrome(num)
	expected := true
	if !answer == expected {
		t.Errorf("Answer should be %v for %v but %v was returned...", expected, num, answer)
	}
}

func TestIsPalindromeCaseFour(t *testing.T) {
	num := 1000021
	answer := isPalindrome(num)
	expected := false
	if !answer == expected {
		t.Errorf("Answer should be %v for %v but %v was returned...", expected, num, answer)
	}
}
