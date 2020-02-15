package parser

import (
	"errors"
	"github.com/Loptt/calculator-interpreter/lexer"
	"strconv"
)

var correct = true

func ParseS(tokens []lexer.Token) (int64, error) {
	index := 0
	var result int64

	match(tokens, lexer.Number, index)
	a, err := strconv.ParseInt(tokens[index].Value(), 10, 64)
	if err != nil {
		return 0, errors.New("Unable to parse int")
	}

	result = a
	index++

	switch tokens[index].Type() {
	case lexer.Plus:
		match(tokens, lexer.Plus, index)
		index++
		match(tokens, lexer.Number, index)
		b, err := strconv.ParseInt(tokens[index].Value(), 10, 64)
		if err != nil {
			return 0, errors.New("Unable to parse int")
		}

		result += b
	case lexer.Minus:
		match(tokens, lexer.Minus, index)
		index++
		match(tokens, lexer.Number, index)
		b, err := strconv.ParseInt(tokens[index].Value(), 10, 64)
		if err != nil {
			return 0, errors.New("Unable to parse int")
		}

		result -= b
	}

	return result, nil
}

func match(tokens []lexer.Token, tt lexer.TokenType, index int) {
	if index >= len(tokens) {
		correct = false
	}

	if tokens[index].Type() != tt {
		correct = false
	}
}
