package compiler

import (
	"bufio"
	"fmt"
	"github.com/praveensanap/glox-interpreter/scanner"
	"io/ioutil"
	"os"
)

func Compile(file string) {
	if file != "" {
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
	scanne.ScanTokens()
}
