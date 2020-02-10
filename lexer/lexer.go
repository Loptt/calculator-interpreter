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

var transMatrix [][]int = [][]int{
	{1, 2, 3, 0},
	{1, 101, 101, 101},
	{102, 102, 102, 102},
	{103, 103, 103, 103},
}

func scan(s string) []Token {
	//index := 0
	return []Token{}
}

func filter(c rune) int {
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
