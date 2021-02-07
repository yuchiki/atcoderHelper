package ach

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/test"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/version"
)

// NewAchCmd returns ach command.
func NewAchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ach",
		Short: "ach automates routine work you does when you participate AtCoder contests",
		Long:  `ach automates routine work you does when you participate AtCoder contests. `,
	}

	registerSubcommands(cmd)

	return cmd
}

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(version.NewVersionCmd())
	cmd.AddCommand(contest.NewContestCmd())
	cmd.AddCommand(test.NewTestCmd())
}

/*
 func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
*/
