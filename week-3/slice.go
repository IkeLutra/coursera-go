package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	values := make([]int, 3)
	var entry string
	for {
		fmt.Printf("Please enter an integer: ")
		fmt.Scan(&entry)
		if entry == "X" {
			break
		}
		value, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Printf("This is not an integer: %v\n", entry)
			continue
		}
		values = append(values, value)

		sort.Ints(values)
		fmt.Printf("Your slice: %v\n", values)
	}

}
