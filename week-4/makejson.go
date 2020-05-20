package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var name, address string
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a name:\n")
	name, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("The name could not be read\n")
		os.Exit(1)
	}
	name = strings.TrimSpace(name)
	fmt.Printf("Please enter an address:\n")
	address, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("The address could not be read\n")
		os.Exit(1)
	}
	address = strings.TrimSpace(address)
	person := Person{
		Name:    name,
		Address: address,
	}
	json, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Could not marshall JSON\n")
		os.Exit(1)
	}
	fmt.Println(string(json))

}
