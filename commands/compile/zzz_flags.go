package compile

import "github.com/spf13/cobra"

const (
	sourceFileFlag = "file"
)

func addSourceFileFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(sourceFileFlag, "f", "test.lox", "file path of the source file")
}

func getSourceFileFlagValue(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(sourceFileFlag)
}
