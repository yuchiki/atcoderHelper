package contest

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest/create"
)

func NewContestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contest",
		Short: "manipulates an AtCoder contest",
		Long:  `manipulates an AtCoder contest.`,
	}

	registerSubcommands(cmd)

	return cmd
}

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(create.NewContestCreateCmd())
}
