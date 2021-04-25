package incoming

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/repository"
	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

var errUnknownInRepositoryLayer = errors.New("an error thrown in fetcher function")

func TestIncoming_Execute(t *testing.T) {
	generateBuilder := func(infos []repository.ContestInfo, err error) func() *cobra.Command {
		return func() *cobra.Command {
			return NewContestIncomingCmd(func() ([]repository.ContestInfo, error) {
				return infos, err
			})
		}
	}

	testutil.TestCaseTemplates{
		testutil.HasName("incoming"),
	}.Build(generateBuilder(nil, nil)).Run(t)

	testutil.TestCases{
		{
			Name: "OK",
			CmdBuilder: generateBuilder([]repository.ContestInfo{
				{ID: "id1", Name: "name1"},
				{ID: "id2", Name: "name2"},
				{ID: "id3", Name: "name3"},
			}, nil),
			OutputCheck: testutil.OutputShouldBe("id1: name1\nid2: name2\nid3: name3\n"),
		},
		{
			Name:       "returns an error when the repository layer fails",
			CmdBuilder: generateBuilder(nil, errUnknownInRepositoryLayer),
			ErrorCheck: testutil.AnyError(),
		},
	}.Run(t)
}
