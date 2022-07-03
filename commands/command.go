package command

import (
	"github.com/praveensanap/glox-interpreter/commands/clibag"
	compileCmd "github.com/praveensanap/glox-interpreter/commands/compile"
	"github.com/praveensanap/glox-interpreter/logger"
	"github.com/spf13/cobra"
	"os"
)

var Version = "0.0.1"

func Run() {
	b := clibag.NewBag()
	rootCmd := cobra.Command{
		Use:     "loxc",
		Short:   "lox compiler",
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Change log level according to flag
			var logLevel logger.Level = logger.DebugLevel
			if cmd.Name() == "__complete" || cmd.Name() == "completion" {
				logLevel = logger.ErrorLevel
			} else if l, err := getLogLevelFlagValue(cmd); err == nil {
				logLevel = logger.LevelFromString(l)
			}
			logger.SetLevel(logLevel)
			return nil
		},
	}

	addLogLevelFlag(&rootCmd)

	b.BuildSubCommands(&rootCmd, compileCmd.NewCmd)

	err := rootCmd.Execute()
	if err != nil {
		rootCmd.PrintErrf("Original error: %+v", err)
		os.Exit(1)
	}
}
