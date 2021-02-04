package version

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yuchiki/atcoderHelper/internal/testutil"
)

func TestVersion_Execute(t *testing.T) {
	jsonFieldsString := []string{
		`"Version":"given by LDFLAGS"`,
		`"Commit":"given by LDFLAGS"`,
		`"Edited":"given by LDFLAGS"`,
		`"Date":"given by LDFLAGS"`,
	}
	expected := fmt.Sprintf("{%s}\n", strings.Join(jsonFieldsString, ","))

	testutil.TestCaseTemplates{
		{
			Name:        "version shows the version",
			OutputCheck: testutil.OutputShouldBe(expected),
		},
	}.
		Build(NewVersionCmd).
		Run(t)
}
