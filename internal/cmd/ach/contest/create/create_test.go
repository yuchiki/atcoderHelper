package create

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestCreate_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.HasName("create"),
		// This is a temporal implementation with temporal args, so the test are given later.
	}.
		Build(NewContestCreateCmd).
		Run(t)
}
