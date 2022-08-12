package compile

import (
	"github.com/praveensanap/glox-interpreter/commands/clibag"
	"github.com/praveensanap/glox-interpreter/compiler"
	"github.com/praveensanap/glox-interpreter/logger"
	"github.com/spf13/cobra"
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

			compiler.Compile(file)
			return nil
		},
	}
	addSourceFileFlag(cmd)
	return cmd
}
