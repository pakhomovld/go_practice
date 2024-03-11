package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Can't read input. Please try again", err)
		return
	}

	input = strings.TrimSpace(input)
	inputLower := strings.ToLower(input)

	if strings.HasPrefix(inputLower, "i") && strings.HasSuffix(inputLower, "n") && strings.Contains(inputLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
