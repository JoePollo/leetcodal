package main

import (
	"testing"
)

func TestRomanToIntsCaseOne(t *testing.T) {
	romanInt := "III"
	answer := romanToInts(romanInt)
	expected := 3
	if answer != expected {
		t.Errorf("Answer should have been %v but %v was returned.", expected, answer)
	}
}

func TestRomanToIntsCaseTwo(t *testing.T) {
	romanInt := "LVIII"
	answer := romanToInts(romanInt)
	expected := 58
	if answer != expected {
		t.Errorf("Answer should have been %v but %v was returned.", expected, answer)
	}
}

func TestRomanToIntsCaseThree(t *testing.T) {
	romanInt := "MCMXCIV"
	answer := romanToInts(romanInt)
	expected := 1994
	if answer != expected {
		t.Errorf("Answer should have been %v but %v was returned.", expected, answer)
	}
}

func TestRomanToIntsCaseFour(t *testing.T) {
	romanInt := "IV"
	answer := romanToInts(romanInt)
	expected := 4
	if answer != expected {
		t.Errorf("Answer should have been %v but %v was returned.", expected, answer)
	}
}

func TestRomanToIntsCaseFive(t *testing.T) {
	romanInt := "XC"
	answer := romanToInts(romanInt)
	expected := 90
	if answer != expected {
		t.Errorf("Answer should have been %v but %v was returned.", expected, answer)
	}
}
