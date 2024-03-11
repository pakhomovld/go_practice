package main

import (
		"fmt"
		"sort"
		"strconv"
)

func main() {
	var input string
	intSlice := make([]int, 0, 3)

	for {
		fmt.Println("Enter an integer or 'X' to exit:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		
		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			continue
		}

		intSlice = append(intSlice, num)

		sort.Ints(intSlice)

		fmt.Println("Sorted slice:", intSlice)
	}
}
