package compiler

import (
	"bufio"
	"fmt"
	"github.com/praveensanap/glox-interpreter/ast"
	"github.com/praveensanap/glox-interpreter/errors"
	"github.com/praveensanap/glox-interpreter/parser"
	"github.com/praveensanap/glox-interpreter/scanner"
	"os"
)

func Compile(file string) {
	if file != "" {
		fmt.Printf("compiling %s\n", file)
		b, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		run(string(b))
	} else {
		runPrompt()
	}

}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || line == nil {
			break
		}
		run(string(line))
	}
}

// compiles a string
func run(s string) {

	// init a scanner.
	scanne := scanner.New(s)

	// scan all tokens
	tokens := scanne.ScanTokens()
	parse := parser.NewParser(tokens)
	expr, err := parse.Parse()
	if err != nil {
		errors.Error(0, err.Error())
	}
	p := ast.Printer{}
	if expr != nil {
		p.Print(expr)
	}
}
