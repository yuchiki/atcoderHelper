package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
)

func TestVersion_Execute(t *testing.T) {
	cmd := newVersionCmd()
	jsonFieldsString := []string{
		`"Version":"given by LDFLAGS"`,
		`"Commit ":"given by LDFLAGS"`,
		`"Edited":"given by LDFLAGS"`,
		`"Date":"given by LDFLAGS"`,
	}
	checkCommand(
		t,
		cmd,
		outputShouldBe(fmt.Sprintf("{%s}\n", strings.Join(jsonFieldsString, ","))),
		shouldNotHaveError(),
	)
}

type (
	OutputCheck func(t *testing.T, output string)
	ErrorCheck  func(t *testing.T, err error)
)

func outputShouldBe(expected string) OutputCheck {
	return func(t *testing.T, output string) {
		t.Helper()

		if diff := cmp.Diff(output, expected); diff != "" {
			t.Error(diff)
		}
	}
}

// nolint:deadcode,unused // 後々必ず使うはず
func errorShouldBe(expected error) ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()

		if !errors.Is(expected, err) {
			t.Errorf("expected error %v, but actual error %v", expected, err)
		}
	}
}

func shouldNotHaveError() ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Error(err)
		}
	}
}

func checkCommand(t *testing.T, command *cobra.Command, outputCheck OutputCheck, errorCheck ErrorCheck) {
	t.Helper()

	var buf bytes.Buffer

	command.SetOut(&buf)
	err := command.Execute()
	errorCheck(t, err)

	output := buf.String()
	outputCheck(t, output)
}
