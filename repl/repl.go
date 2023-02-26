package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }

		p := parser.New(l)
		program := p.ParseProgram()
		errors := p.Errors()
		if len(errors) > 0 {
			fmt.Fprintf(out, "parser has %d errors\n", len(errors))
			for _, msg := range errors {
				fmt.Fprintf(out, "parser errors :%q\n", msg)
			}
		} else {
			for i := range program.Statements {
				stmt := program.Statements[i]
				fmt.Printf("Statement %d is {%q}.\n", i, stmt.TokenLiteral())
			}
		}
	}
}
