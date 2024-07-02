package repl

import (
	"bufio"
	"fmt"
	"io"
	"go-lexer/lexer"
    "go-lexer/token"
)
const PROMPT = ">> "

func Start(in io.Reader, ou io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
