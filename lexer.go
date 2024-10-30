package main

import (
	"strings"
	"unicode"
)

// Token types
const (
	INT    = "INT"     // Integer token
	PLUS   = "PLUS"    // "+" token
	MINUS  = "MINUS"   // "-" token
	EOF    = "EOF"     // End of File/Line token
	UNKNOWN = "UNKNOWN" // Unknown token
)

// Token structure
type Token struct {
	Type    string
	Literal string
}

// Lexer function: converts input into tokens
func Lexer(input string) []Token {
	var tokens []Token
	runes := []rune(strings.TrimSpace(input)) // Convert input to runes for better handling

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if unicode.IsDigit(char) {
			num := string(char)
			// Handle multi-digit numbers
			for i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				i++
				num += string(runes[i])
			}
			tokens = append(tokens, Token{Type: INT, Literal: num})
		} else if char == '+' {
			tokens = append(tokens, Token{Type: PLUS, Literal: string(char)})
		} else if char == '-' {
			tokens = append(tokens, Token{Type: MINUS, Literal: string(char)})
		} else {
			tokens = append(tokens, Token{Type: UNKNOWN, Literal: string(char)})
		}
	}

	// Append EOF token
	tokens = append(tokens, Token{Type: EOF, Literal: ""})
	return tokens
}
