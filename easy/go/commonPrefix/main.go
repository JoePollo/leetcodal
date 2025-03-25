package main

import "fmt"

func LongestCommonPrefix(strings []string) string {
	commonPrefix := ""
	newPrefix := ""
	nextWord := false

	for index, word := range strings {
		if nextWord {
			continue
		}
		if index == 0 && word[0] == strings[1][0] {
			newPrefix += string(word[0])
			continue
		} else if newPrefix != "" {
			for _index, letter := range word[1:] {
				if index+1 < len(strings) && string(letter) == string(strings[index+1][_index]) {
					newPrefix += string(letter)
					continue
				}
			}
		} else {
			if commonPrefix != "" {
				if len(newPrefix) > len(commonPrefix) {
					commonPrefix = newPrefix
				}
			} else {
				commonPrefix = newPrefix
			}
		}
	}

	return commonPrefix
}

func main() {
	fmt.Print("tacos")
}
