package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortSubarray(part []int, ch chan []int) {
	fmt.Println("Sorting subarray:", part)
	sort.Ints(part)
	ch <- part
}

func mergeTwoArrays(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main() {
	fmt.Println("Enter integers separated by space:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strs := strings.Fields(input)

	nums := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid input:", s)
			return
		}
		nums[i] = n
	}

	if len(nums) < 4 {
		fmt.Println("Please enter at least 4 integers.")
		return
	}

	// Your original partition logic using int division
	partSize := int(math.Round(float64(len(nums)) / 4.0))
	ch := make(chan []int, 4)

	// Explicitly defined ranges for 4 partitions
	start1 := 0
	end1 := start1 + partSize

	start2 := end1
	end2 := start2 + partSize

	start3 := end2
	end3 := start3 + partSize

	start4 := end3
	end4 := len(nums) // last partition takes the remainder

	// Explicit goroutines (not in a loop)
	go sortSubarray(nums[start1:end1], ch)
	go sortSubarray(nums[start2:end2], ch)
	go sortSubarray(nums[start3:end3], ch)
	go sortSubarray(nums[start4:end4], ch)

	// Collect sorted partitions
	sortedParts := make([][]int, 0, 4)
	for range 4 {
		sortedParts = append(sortedParts, <-ch)
	}

	// Merge all sorted parts
	merged := mergeTwoArrays(sortedParts[0], sortedParts[1])
	merged = mergeTwoArrays(merged, sortedParts[2])
	merged = mergeTwoArrays(merged, sortedParts[3])

	fmt.Println("Fully sorted array:", merged)
}
