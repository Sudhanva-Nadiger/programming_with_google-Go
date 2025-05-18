package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// There is a function called Swap(), which takes two arguments -- a slice of integers, and an index value i.
func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func BubbleSort(nums []int) {
	for i := range nums {
		for j := range len(nums) - 1 - i {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func main() {
	fmt.Println("Enter up to 10 integers separated by spaces:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Split and parse input
	parts := strings.Fields(input)
	if len(parts) > 10 {
		fmt.Println("You entered more than 10 numbers. Only the first 10 will be used.")
		parts = parts[:10]
	}

	var numbers []int
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			fmt.Printf("Invalid input: %s is not an integer\n", p)
			return
		}
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)

	fmt.Println("Sorted numbers:", numbers)
}
