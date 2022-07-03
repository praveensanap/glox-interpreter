package compile

import (
	"bufio"
	"fmt"
	"github.com/praveensanap/glox-interpreter/commands/clibag"
	"github.com/praveensanap/glox-interpreter/logger"
	"github.com/praveensanap/glox-interpreter/scanner"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func NewCmd(b *clibag.Bag) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compile [flags]",
		Short: "Compile from source",
		//Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := getSourceFileFlagValue(cmd)
			if err != nil {
				logger.Error("source file required")
			}

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
			return nil
		},
	}
	addSourceFileFlag(cmd)
	return cmd
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
	fmt.Println(s)
}
