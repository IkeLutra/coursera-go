package main

import (
	"fmt"
	"math"
	"os"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	var err error
	fmt.Print("Acceleration: ")
	_, err = fmt.Scan(&acceleration)
	if err != nil {
		fmt.Printf("Argument invalid: %v", err)
		os.Exit(1)
	}
	fmt.Print("Initial Velocity: ")
	_, err = fmt.Scan(&initialVelocity)
	if err != nil {
		fmt.Printf("Argument invalid: %v", err)
		os.Exit(1)
	}
	fmt.Print("Initial Displacement: ")
	_, err = fmt.Scan(&initialDisplacement)
	if err != nil {
		fmt.Printf("Argument invalid: %v", err)
		os.Exit(1)
	}
	displaceFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Print("Time: ")
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Printf("Argument invalid: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Displacement: %v", displaceFn(time))
}
