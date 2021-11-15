package contest

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestContest_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.HasName("contest"),
		testutil.HasSubcommands("create", "incoming", "recent", "testcase"),
	}.
		Build(NewContestCmd).
		Run(t)
}
