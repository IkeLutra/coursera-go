package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(elements []int, index int) {
	a := elements[index]
	b := elements[index+1]
	elements[index] = b
	elements[index+1] = a
}

func BubbleSort(elements []int) {
	n := len(elements) - 1
	for {
		swapped := false
		for i := 0; i < n; i++ {
			if elements[i] > elements[i+1] {
				Swap(elements, i)
				swapped = true
			}
		}
		n--
		if !swapped {
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num string
	var err error
	var elements = []int{}

	fmt.Print("Please enter up to 10 numbers with a newline between each. Press X to finish the list early\n")
	for {
		num, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("That number could not be read\n")
			os.Exit(1)
		} else {
			num = strings.ToUpper(strings.TrimSpace(num))
			if num == "X" {
				break
			}
			element, err := strconv.Atoi(num)
			if err == nil {
				elements = append(elements, element)
			}

		}
		if len(elements) >= 10 {
			break
		}
	}
	BubbleSort(elements)
	fmt.Printf("Sorted: %v", elements)
}
