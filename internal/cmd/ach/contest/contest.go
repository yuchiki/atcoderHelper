package contest

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest/create"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest/incoming"
	"github.com/yuchiki/atcoderHelper/internal/repository"
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
	cmd.AddCommand(incoming.NewContestIncomingCmd(repository.FetchIncoming))
}
