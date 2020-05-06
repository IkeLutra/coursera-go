package main

import "fmt"

func main() {
	var entry float64
	fmt.Printf("Please enter a floating point number: ")
	fmt.Scan(&entry)
	fmt.Printf("Your truncated number: %d\n", int(entry))
}
