package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, address string
	var err error
	person := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a name:\n")
	name, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("The name could not be read\n")
		os.Exit(1)
	}
	name = strings.TrimSpace(name)
	person["name"] = name
	fmt.Printf("Please enter an address:\n")
	address, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("The address could not be read\n")
		os.Exit(1)
	}
	address = strings.TrimSpace(address)
	person["address"] = address
	json, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Could not marshall JSON\n")
		os.Exit(1)
	}
	fmt.Println(string(json))

}
