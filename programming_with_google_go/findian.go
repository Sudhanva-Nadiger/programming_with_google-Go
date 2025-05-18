package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mains() {
	// Prompt user for input
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Clean up and normalize the input
	input = strings.TrimSpace(input)
	inputLower := strings.ToLower(input)

	// Check conditions
	startsWithI := strings.HasPrefix(inputLower, "i")
	endsWithN := strings.HasSuffix(inputLower, "n")
	containsA := strings.Contains(inputLower, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
