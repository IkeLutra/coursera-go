package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	type test struct {
		input  []int
		output [][]int
	}

	tests := []test{
		{input: []int{76, 89, 18, 34, 98, 97, 92, 80, 83, 16}, output: [][]int{[]int{76, 89, 18}, []int{34, 98, 97}, []int{92, 80}, []int{83, 16}}},
		{input: []int{76, 89, 18, 34, 98, 97, 92, 80}, output: [][]int{[]int{76, 89}, []int{18, 34}, []int{98, 97}, []int{92, 80}}},
		{input: []int{76, 89, 18, 34, 98, 97, 92}, output: [][]int{[]int{76, 89}, []int{18, 34}, []int{98, 97}, []int{92}}},
	}

	for _, tc := range tests {
		output := sliceIntoFour(tc.input)
		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("Input %v Output %v does not match %v", tc.input, tc.output, output)
		}
	}

}

func TestSortSlice(t *testing.T) {
	inputSlice := []int{76, 89, 18, 34}
	sortSlice(inputSlice)
	fmt.Printf("Input Slice After: %v", inputSlice)
	if !reflect.DeepEqual(inputSlice, []int{18, 34, 76, 89}) {
		t.Errorf("Slice not sorted: %v", inputSlice)
	}
}
