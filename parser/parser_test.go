package parser

import (
	"testing"
	"github.com/Loptt/calculator-interpreter/lexer"
)

func TestParseS(t *testing.T) {
	tests := []struct {
		in []lexer.Token
		want int64
	}{
		{
			in: []lexer.Token{
				lexer.NewToken(lexer.Number, "4"),
				lexer.NewToken(lexer.Plus, "+"),
				lexer.NewToken(lexer.Number, "5"),				
			},
			want: 9,
		},
		{
			in: []lexer.Token{
				lexer.NewToken(lexer.Number, "11"),
				lexer.NewToken(lexer.Minus, "-"),
				lexer.NewToken(lexer.Number, "6"),				
			},
			want: 5,
		},
		{
			in: []lexer.Token{
				lexer.NewToken(lexer.Number, "234"),
				lexer.NewToken(lexer.Plus, "+"),
				lexer.NewToken(lexer.Number, "67543"),				
			},
			want: 67777,
		},
	}

	for _, test := range tests {
		got, err := ParseS(test.in)
		if err != nil {
			t.Errorf("ParseS(%v) Error: %v", test.in, err)
		}
		if got != test.want {
			t.Errorf("ParseS(%v) = %v, expected %v", test.in, got, test.want)
		}
	}
}