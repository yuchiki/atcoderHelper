package ach

import (
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestAch_Execute(t *testing.T) {
	testutil.TestCaseTemplates{
		testutil.HasName("ach"),
		testutil.HasSubcommands("version", "contest"),
	}.
		Build(NewAchCmd).
		Run(t)
}
