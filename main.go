package main

import "fmt"

func main() {
	input := "5 + 3 - 2"
	fmt.Println("Input:", input)


	tokens := Lexer(input)
	fmt.Println("Tokens:", tokens)


	result := Parse(tokens)
	fmt.Println("Result:", result)
}
