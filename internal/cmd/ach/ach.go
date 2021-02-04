package ach

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/version"
)

func NewAchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ach",
		Short: "ach は Atcoder出場に際する定形作業を自動化します",
		Long:  `ach は Atcoder出場に際する定形作業を自動化します。`,
	}

	registerSubcommands(cmd)

	return cmd
}

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(version.NewVersionCmd())
	cmd.AddCommand(contest.NewContestCmd())
}

/*
 func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
*/
