package lexer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		in   byte
		want int
	}{
		{'1', 0},
		{'5', 0},
		{'+', 1},
		{'-', 2},
		{'9', 0},
		{'0', 0},
		{' ', 3},
		{'\t', 3},
	}

	for _, test := range tests {
		got := filter(test.in)
		if got != test.want {
			t.Errorf("Filter(%c) = %d, expected %d", test.in, got, test.want)
		}
	}
}

func TestScan(t *testing.T) {
	tests := []struct {
		in   string
		want []Token
	}{
		{
			in: "2 3",
			want: []Token{
				Token{Number, "2"},
				Token{Number, "3"},
			},
		},
		{
			in: "1+4",
			want: []Token{
				Token{Number, "1"},
				Token{Plus, "+"},
				Token{Number, "4"},
			},
		},
		{
			in: "12-344  ",
			want: []Token{
				Token{Number, "12"},
				Token{Minus, "-"},
				Token{Number, "344"},
			},
		},
		{
			in: "     5 + 6 ",
			want: []Token{
				Token{Number, "5"},
				Token{Plus, "+"},
				Token{Number, "6"},
			},
		},
	}

	for _, test := range tests {
		got := Scan(test.in)
		if !cmp.Equal(got, test.want) {
			t.Errorf("Scan(%v) = %v, expected %v", test.in, got, test.want)
		}
	}
}
