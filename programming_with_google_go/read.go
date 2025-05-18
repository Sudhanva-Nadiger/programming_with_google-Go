package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FullName struct {
	Fname [20]byte
	Lname [20]byte
}

func mainssss() {
	// Prompt the user for the filename
	fmt.Print("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []FullName
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			fmt.Println("Skipping malformed line:", line)
			continue
		}

		var n FullName
		copy(n.Fname[:], parts[0])
		copy(n.Lname[:], parts[1])
		names = append(names, n)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print out the collected names
	fmt.Println("\nNames read from file:")
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", string(n.Fname[:]), string(n.Lname[:]))
	}
}
