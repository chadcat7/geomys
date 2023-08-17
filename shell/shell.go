package shell

import (
	"bufio"
	"fmt"
	"geomys/lexer"
	"geomys/token"
	"io"
)

const PROMPT = "geomys > "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLex(line)
		for tok := l.AdvanceToken(); tok.Type != token.EOF; tok = l.AdvanceToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
