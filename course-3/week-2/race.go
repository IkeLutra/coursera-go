package main

import (
	"fmt"
	"time"
)

// Explanation
// Both go routines modify the same variable. As order in which the loops are
// run is not guranteed by go, the printed listed of values will differ each
// time the program is run. This is the race condition.
// For example all of the additions could happen first resulting in the program printing 0 to 9 and then 9 to 0
// Or the subtractions could all happen first printing in 0 to -9 and then -9 to 0
// Or the additions and subtractions could be interleaved in an unpredictable order resulting in a different list printing

func subtract(count *int) {
	for i := 0; i < 10; i++ {
		*count = *count - 1
		fmt.Printf("Sub: %v\n", *count)
	}
}

func add(count *int) {
	for i := 0; i < 10; i++ {
		*count = *count + 1
		fmt.Printf("Add: %v\n", *count)
	}
}

func main() {
	count := 0
	go subtract(&count)
	go add(&count)
	time.Sleep(time.Second)
}
