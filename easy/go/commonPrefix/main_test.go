package main

import "testing"

func TestLongestCommonPrefixCaseOne(t *testing.T) {
	strings := []string{
		"flower",
		"flow",
		"flight",
	}
	answer := LongestCommonPrefix(strings)
	expected := "fl"
	if answer != expected {
		t.Errorf("Answer should be %v but %v was returned...", expected, answer)
	}
}
