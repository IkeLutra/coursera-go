package main

import (
	"testing"
)

func TestDisplaceFn(t *testing.T) {
	type test struct {
		time         float64
		displacement float64
	}
	tests := []test{
		{time: 3, displacement: 52},
		{time: 5, displacement: 136},
	}
	displaceFn := GenDisplaceFn(10, 2, 1)
	for _, tc := range tests {
		resultDisplacement := displaceFn(tc.time)
		if resultDisplacement != tc.displacement {
			t.Errorf("Could not calculate displacement %v != %v", resultDisplacement, tc.displacement)
		}
	}
}
