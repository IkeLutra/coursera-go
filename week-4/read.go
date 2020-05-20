package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Forename string
	Surname  string
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal(errors.New("Please supply a file of names"))
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	people := make([]Person, 0)
	for scanner.Scan() {
		nameParts := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(nameParts) != 2 {
			fmt.Println("Invalid name skipping")
			continue
		}
		people = append(people, Person{
			Forename: nameParts[0],
			Surname:  nameParts[1],
		})
	}
	for _, person := range people {
		fmt.Printf("Forename: %v Surname: %v\n", person.Forename, person.Surname)
	}

}
