package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

func NewCow() *cow {
	return &cow{
		"grass",
		"walk",
		"moo",
	}
}

type cow struct {
	food       string
	locomotion string
	noise      string
}

func (a *cow) Eat() {
	fmt.Println(a.food)
}

func (a *cow) Move() {
	fmt.Println(a.locomotion)
}
func (a *cow) Speak() {
	fmt.Println(a.noise)
}

func NewBird() *bird {
	return &bird{
		"worms",
		"fly",
		"peep",
	}
}

type bird struct {
	food       string
	locomotion string
	noise      string
}

func (a *bird) Eat() {
	fmt.Println(a.food)
}

func (a *bird) Move() {
	fmt.Println(a.locomotion)
}
func (a *bird) Speak() {
	fmt.Println(a.noise)
}

func NewSnake() *snake {
	return &snake{
		"mice",
		"slither",
		"hsss",
	}
}

type snake struct {
	food       string
	locomotion string
	noise      string
}

func (a *snake) Eat() {
	fmt.Println(a.food)
}

func (a *snake) Move() {
	fmt.Println(a.locomotion)
}
func (a *snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var animals map[string]Animal
	animals = make(map[string]Animal)
	var name, command, actionOrType string
	for {
		fmt.Print("> ")
		fmt.Scan(&command, &name, &actionOrType)
		command = strings.ToLower(command)
		actionOrType = strings.ToLower(actionOrType)

		if command == "newanimal" {
			var animal Animal
			switch actionOrType {
			case "cow":
				animal = NewCow()
			case "bird":
				animal = NewBird()
			case "snake":
				animal = NewSnake()
			}
			if animal != nil {
				animals[name] = animal
				fmt.Println("Created it!")
			} else {
				fmt.Printf("Could not create animal type: %v\n", actionOrType)
			}
		} else if command == "query" {
			animal, ok := animals[name]
			if ok {
				switch actionOrType {
				case "eat":
					animal.Eat()
				case "speak":
					animal.Speak()
				case "move":
					animal.Move()
				default:
					fmt.Printf("Invalid action %v\n", actionOrType)
				}
			} else {
				fmt.Printf("Could not find animal called: %v\n", name)
			}
		} else {
			fmt.Printf("Invalid command: %s\n", command)
		}
	}
}
