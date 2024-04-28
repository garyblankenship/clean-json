package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func findLargestJSON(input string) string {
	var stack []int
	largestJSON := ""
	maxLength := 0

	// Iterate over each character in the input string
	for i, char := range input {
		switch char {
		case '{', '[': // Handle JSON objects and arrays
			stack = append(stack, i) // Push the index of the opening brace or bracket
		case '}', ']': // Handle closing of JSON objects and arrays
			if len(stack) == 0 {
				continue // Avoid underflow if the stack is empty
			}
			start := stack[len(stack)-1] // Get the last unmatched opening brace or bracket
			stack = stack[:len(stack)-1] // Pop from stack

			// Extract the substring from start to current position
			substr := input[start : i+1]
			var js json.RawMessage
			if json.Unmarshal([]byte(substr), &js) == nil {
				// Check if this is the largest valid JSON found
				if len(substr) > maxLength {
					maxLength = len(substr)
					largestJSON = substr
				}
			}
		}
	}

	return largestJSON // Return the largest valid JSON, or an empty string if none found
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input strings.Builder

	// Read input from stdin
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1) // Exit on errors other than EOF
		}
		input.WriteString(line)
		if err == io.EOF {
			break // Stop reading on EOF
		}
	}

	// Process the input to find the largest valid JSON
	result := findLargestJSON(input.String())
	if result == "" {
		fmt.Println("No valid JSON found")
	} else {
		fmt.Println(result)
	}
}
