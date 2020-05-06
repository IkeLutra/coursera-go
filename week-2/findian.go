package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var entry string
	fmt.Printf("Please enter a string to search:\n")
	reader := bufio.NewReader(os.Stdin)
	entry, err := reader.ReadString('\n')
	if err == nil {
		entry = strings.TrimSpace(strings.ToLower(entry))
		if strings.HasPrefix(entry, "i") && strings.HasSuffix(entry, "n") && strings.Contains(entry, "a") {
			fmt.Printf("Found!\n")
		} else {
			fmt.Printf("Not Found!\n")
		}
	} else {
		fmt.Print("An error occurred")
	}

}
