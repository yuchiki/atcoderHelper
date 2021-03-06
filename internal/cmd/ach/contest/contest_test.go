package contest

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestAch_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.HasName("contest"),
		testutil.HasSubcommands("create"),
		testutil.HasSubcommands("incoming"),
	}.
		Build(NewContestCmd).
		Run(t)
}
