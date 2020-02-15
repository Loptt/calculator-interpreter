package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/Loptt/calculator-interpreter/parser"
	"github.com/Loptt/calculator-interpreter/lexer"
)

func main()  {
	var input string
	reader := bufio.NewReader(os.Stdin)

	for true {
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		tokens := lexer.Scan(input)
		result, err := parser.ParseS(tokens)
		if err != nil {
			fmt.Errorf("Error: %v", err)
		} else {
			fmt.Println(result)
		}
	}
}