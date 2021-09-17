package testcase

import (
	"github.com/spf13/cobra"

	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest/testcase/fetch"
	"github.com/yuchiki/atcoderHelper/internal/repository"
)

func NewContestTestcaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testcase",
		Short: "manipulates testcases",
		Long:  `manipulates testcases.`,
	}

	registerSubcommands(cmd)

	return cmd
}

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(fetch.NewTestcaseFetchCmd(repository.FetchTestcases))
}
