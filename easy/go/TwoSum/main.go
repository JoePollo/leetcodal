package main

import (
	"fmt"
)

func twoSum(num []int, target int) ([]int, error) {
	target_indices := []int{}
	num_length := len(num)
	for index, number := range num {
		for i := range num_length {
			if number+num[index+i] == target {
				target_indices = append(target_indices, index, index+i)
				return target_indices, nil
			} else {
				continue
			}
		}
	}
	return nil, fmt.Errorf("Matching indices were not found.")
}

func main() {
	target_indices, _ := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(target_indices)
}
