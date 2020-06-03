package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	numbers := []int{3, 5, 1}
	Swap(numbers, 1)
	if numbers[1] != 1 && numbers[2] != 5 {
		t.Errorf("Did not swap: %v", numbers)
	}
}

func TestBubbleSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := []int{7, 8, 2, 1, 3, 4, 5, 10, 6, 9}
	BubbleSort(actual)
	for i := range sorted {
		if sorted[i] != actual[i] {
			t.Errorf("Index %d not equal %d %d", i, actual[i], sorted[i])
		}
	}

	sorted = []int{-85, -54, -52, -37, -35, -27, -17, -11, 15, 17, 42, 45, 56, 60, 81, 86}
	actual = []int{-35, -85, -52, -17, 86, -54, 56, 60, 15, -37, -27, 42, -11, 81, 17, 45}
	BubbleSort(actual)
	for i := range sorted {
		if sorted[i] != actual[i] {
			t.Errorf("Index %d not equal %d %d", i, actual[i], sorted[i])
		}
	}
}
