package main

import "fmt"

func main() {
	// Sample input to our language
	input := "5 + 3 - 2"
	fmt.Println("Input:", input)

	// Step 1: Lexing
	tokens := Lexer(input)
	fmt.Println("Tokens:", tokens)

	// Step 2: Parsing and Evaluating
	result := Parse(tokens)
	fmt.Println("Result:", result)
}
