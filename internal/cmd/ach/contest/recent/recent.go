package recent

import (
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/repository"
)

func NewContestRecentCmd(fetcher func() ([]repository.ContestInfo, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recent",
		Short: "show recent contests",
		Long:  "show recent contests.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			contests, err := fetcher()
			if err != nil {
				return err
			}

			for _, contest := range contests {
				cmd.Printf("%v: %v\n", contest.ID, contest.Name)
			}

			return nil
		},
	}

	return cmd
}
