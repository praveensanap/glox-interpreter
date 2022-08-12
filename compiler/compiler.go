package compiler

import (
	"bufio"
	"fmt"
	"github.com/praveensanap/glox-interpreter/ast"
	"github.com/praveensanap/glox-interpreter/errors"
	"github.com/praveensanap/glox-interpreter/parser"
	"github.com/praveensanap/glox-interpreter/scanner"
	"io/ioutil"
	"os"
)

func Compile(file string) {
	if file != "" {
		fmt.Printf("compiling %s\n", file)
		b, err := ioutil.ReadFile(file)
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

func run(s string) {
	scanne := scanner.New(s)
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
