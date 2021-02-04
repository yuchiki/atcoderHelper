package ach

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestAch_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.RootSucceeds(),
		testutil.HasSubcommand("version"),
	}.
		Build(NewAchCmd).
		Run(t)
}
