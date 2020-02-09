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

const transMatrix [][]int = [][]int{
	{1, 2, 3, 0},
	{1, 101, 101, 101},
	{102, 102, 102, 102},
	{103, 103, 103, 103},
}

func scan(s string) []Token {
	index := 0

}

func filter(c rune) int {
	switch c {
	case '1':
	case '2':
		return 0
	default:
		return 999
	}
}
