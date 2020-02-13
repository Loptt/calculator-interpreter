package lexer

import (
	"strings"
)

// TokenType defines an enum of the possible types of a token
type TokenType int

const (
	Number TokenType = iota + 1
	Plus
	Minus
)

// Token represents a lexeme in the language
type Token struct {
	tt  TokenType
	val string
}

// Equal overrides default token equality
func (t Token) Equal(t2 Token) bool {
	return t.tt == t2.tt && t.val == t2.val
}

var transMatrix [][]int = [][]int{
	{1, 2, 3, 0},
	{1, 101, 101, 101},
	{102, 102, 102, 102},
	{103, 103, 103, 103},
}

// Scan processes an input string and deconstructs it into its corresponding tokens
func Scan(s string) []Token {
	s = s + " "
	index := 0
	state := 0

	var value strings.Builder
	var tokens []Token

	var curr byte

	for index < len(s) {
		for state < 100 && index < len(s) {
			curr = s[index]
			state = transMatrix[state][filter(curr)]
			if state < 100 && curr != ' ' {
				value.WriteByte(curr)
			}
			if state > 100 && curr != ' ' {
				index--
			}
			index++
		}

		switch state {
		case 101:
			tokens = append(tokens, Token{Number, value.String()})
		case 102:
			tokens = append(tokens, Token{Plus, value.String()})
		case 103:
			tokens = append(tokens, Token{Minus, value.String()})
		}

		state = 0
		value.Reset()
	}
	return tokens
}

func filter(c byte) int {
	switch c {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		return 0
	case '+':
		return 1
	case '-':
		return 2
	case ' ', '\t':
		return 3
	}

	return 999
}
