package lexer

import "testing"

func TestFilter(t *testing.T) {
	tests := []struct {
		in   rune
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
