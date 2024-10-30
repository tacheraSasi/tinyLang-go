package main

import "fmt"

func Parse(tokens []Token) int {
	result := 0
	i := 0

	// Assume the first token is an integer
	if tokens[i].Type == INT {
		result = toInt(tokens[i].Literal)
		i++
	}

	// Process remaining tokens
	for i < len(tokens) && tokens[i].Type != EOF {
		token := tokens[i]

		if token.Type == PLUS {
			i++
			result += toInt(tokens[i].Literal)
		} else if token.Type == MINUS {
			i++
			result -= toInt(tokens[i].Literal)
		}
		i++
	}

	return result
}

// Helper function: converts string to integer
func toInt(literal string) int {
	var num int
	fmt.Sscan(literal, &num)
	return num
}
