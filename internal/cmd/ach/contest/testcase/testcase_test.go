package testcase

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestAch_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.HasName("testcase"),
		testutil.HasSubcommands("fetch"),
	}.
		Build(NewContestTestcaseCmd).
		Run(t)
}
