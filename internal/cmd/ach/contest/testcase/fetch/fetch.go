package fetch

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/repository"
)

func NewTestcaseFetchCmd(fetcher func(contest string, task string) ([]repository.Testcase, error)) *cobra.Command {
	contest := new(string)
	task := new(string)

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "fetch testcases",
		Long:  "fetch testcases.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			testcases, err := fetcher(*contest, *task)
			if err != nil {
				return err
			}

			cmd.Print(testcases) // とりあえずdebug出力

			return nil
		},
	}

	cmd.Flags().StringVar(contest, "contest", "", "contestID")
	cmd.Flags().StringVar(task, "task", "", "task")

	return cmd
}
