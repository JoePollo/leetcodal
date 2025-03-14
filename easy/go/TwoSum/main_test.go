package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	num := []int{2, 7, 11, 15}
	target := 9
	target_indices, err := twoSum(num, target)
	answer := []int{0, 1}
	if !reflect.DeepEqual(target_indices, answer) {
		t.Errorf("target_indices value of %v should be %v", target_indices, answer)
	} else if err != nil {
		t.Errorf("err should have value of nil, has value of %v", err)
	}
}
