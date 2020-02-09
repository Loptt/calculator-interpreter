package lexer

type TokenType int

const (
	Number TokenType = iota + 1
	Plus
	Minus
)

type Token struct {
	tt  TokenType
	val string
}

const transMatrix := {{1,2},{3,4}}

func scan(s string) []Token {
	index := 0
	
}