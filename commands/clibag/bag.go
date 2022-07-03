package clibag

import (
	"github.com/spf13/cobra"
)

type Bag struct {
	//Config              *config.Config
	ConfigPath          string
	RefreshRepositories bool
	//OutputFormatter     formatter.OutputFormatter
}

type CommandBuilder func(b *Bag) *cobra.Command

func NewBag() *Bag {
	return &Bag{}
}

func (b *Bag) BuildSubCommands(cmd *cobra.Command, builders ...CommandBuilder) {
	for _, f := range builders {
		cmd.AddCommand(f(b))
	}
}
