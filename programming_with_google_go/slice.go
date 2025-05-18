package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func maisn() {

	arr := make([]int, 3)
	i := 0

	fmt.Println("Instructions: ")
	fmt.Println("Enter the numbers one by one")
	fmt.Println("Press 'X' or 'x' to quit")

	var input string

	for {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "x" {
			fmt.Println("Exiting...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if i < 3 {
			arr[i] = num
			i++
			nums := make([]int, i)
			copy(nums, arr[:i])
			sort.Ints(nums)
			fmt.Println(nums)
		} else {
			arr = append(arr, num)
			sort.Ints(arr)
			fmt.Println("Sorted array:", arr)
		}
	}

}
