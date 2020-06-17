package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	fmt.Println(a.food)
	return a.food
}

func (a *Animal) Move() string {
	fmt.Println(a.locomotion)
	return a.locomotion
}
func (a *Animal) Speak() string {
	fmt.Println(a.noise)
	return a.noise
}

func main() {
	var name, command string
	animals := map[string]Animal{
		"cow": Animal{
			"grass", "walk", "moo",
		},
		"bird": Animal{
			"worms", "fly", "peep",
		},
		"snake": Animal{
			"mice", "slither", "hsss",
		},
	}
	for {
		fmt.Print("> ")
		fmt.Scan(&name, &command)
		name = strings.ToLower(name)
		command = strings.ToLower(command)

		animal, ok := animals[name]

		if ok {
			switch command {
			case "eat":
				animal.Eat()
			case "speak":
				animal.Speak()
			case "move":
				animal.Move()
			default:
				fmt.Printf("Invalid command %v\n", command)
			}
		} else {
			fmt.Printf("Could not find animal: %v\n", name)
		}
	}
}
