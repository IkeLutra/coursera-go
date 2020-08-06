package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sliceIntoFour(input []int) [][]int {
	var length = len(input)
	var sliceSize = length / 4
	var mod = length % 4
	var slices [][]int
	start := 0
	end := sliceSize + 1
	for i := 0; i < 4; i++ {
		end = start + sliceSize
		if i < mod {
			end++
		}
		// fmt.Printf("i: %v, start: %v, end: %v, mod: %v, length: %v, sliceSize: %v\n", i, start, end, mod, length, sliceSize)
		slices = append(slices, input[start:end])
		start = end
	}
	return slices
}

func sortSlice(wg *sync.WaitGroup, input []int) {
	fmt.Printf("Slice to be sorted: %v\n", input)
	sort.Ints(input)
	wg.Done()
}

func main() {
	var input []int
	var inputString string
	fmt.Print("Please enter a list of integers separated by commas: ")
	fmt.Scan(&inputString)
	for _, v := range strings.Split(strings.TrimSuffix(strings.TrimSpace(inputString), ","), ",") {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Printf("Could not convert %v to integer", v)
			os.Exit(1)
		}
		input = append(input, number)
	}
	fmt.Printf("Input: %v\n", input)
	var wg sync.WaitGroup
	slices := sliceIntoFour(input)
	for i, _ := range slices {
		wg.Add(1)
		go sortSlice(&wg, slices[i])
	}
	wg.Wait()
	// You could just use the input again as the slices are just a reference to this
	// but the assignment specifically asked for the slices to be merged back together
	var output []int
	for i, _ := range slices {
		output = append(output, slices[i]...)
	}
	sort.Ints(output)
	fmt.Printf("Sorted array: %v\n", output)
}
