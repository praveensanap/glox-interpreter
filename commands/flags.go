package command

import "github.com/spf13/cobra"

const (
	logLevelFlag = "log"
)

func addLogLevelFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(logLevelFlag, "l", "DEBUG", "set the log `LEVEL`")
	_ = cmd.RegisterFlagCompletionFunc(logLevelFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{
			"DEBUG",
			"INFO",
			"WARN",
			"ERROR",
		}, cobra.ShellCompDirectiveDefault
	})
}

func getLogLevelFlagValue(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(logLevelFlag)
}
