package testutil

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
)

type (
	// OutputCheck is a check for output.
	OutputCheck func(t *testing.T, output string)

	// ErrorCheck is a check for error.
	ErrorCheck func(t *testing.T, err error)

	// TestCase is a tuple for testing.
	TestCase struct {
		Name        string
		CmdBuilder  func() *cobra.Command
		Args        []string
		OutputCheck OutputCheck
		ErrorCheck  ErrorCheck
	}

	// TestCases is a suit of testCases.
	TestCases []TestCase

	// TestCaseTemplate is a variation of TestCase which lacks CmdBuilder.
	TestCaseTemplate struct {
		Name        string
		Args        []string
		OutputCheck OutputCheck
		ErrorCheck  ErrorCheck
	}

	// TestCaseTemplates is a suit of testCaseTemplates.
	TestCaseTemplates []TestCaseTemplate
)

// OutputShouldBe returns a check for output.
func OutputShouldBe(expected string) OutputCheck {
	return func(t *testing.T, output string) {
		t.Helper()

		if diff := cmp.Diff(output, expected); diff != "" {
			t.Error(diff)
		}
	}
}

// ErrorShouldBe returns a check for error.
func ErrorShouldBe(expected error) ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()

		if !errors.Is(expected, err) {
			t.Errorf("expected error %v, but actual error %v", expected, err)
		}
	}
}

// DoesNotCareOutput returns a check which allows everything.
func DoesNotCareOutput() OutputCheck {
	return func(_ *testing.T, _ string) {}
}

// ShouldNotHaveError returns a check which expects no errors.
func ShouldNotHaveError() ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Error(err)
		}
	}
}

// CheckCommand checks if the command behaves expectedly.
func CheckCommand(t *testing.T, command *cobra.Command, outputCheck OutputCheck, errorCheck ErrorCheck, args ...string) {
	t.Helper()

	var buf bytes.Buffer

	command.SetOut(&buf)
	command.SetArgs(args)
	err := command.Execute()
	errorCheck(t, err)

	output := buf.String()
	outputCheck(t, output)
}

// Run runs CheckCommand with this testcase.
// Without OutputCheck, it does not check output.
// WIthout ErrorCheck, it expects no errors.
func (c TestCase) Run(t *testing.T) {
	t.Helper()

	var outputCheck OutputCheck

	if c.OutputCheck == nil {
		outputCheck = DoesNotCareOutput()
	} else {
		outputCheck = c.OutputCheck
	}

	var errorCheck ErrorCheck
	if c.ErrorCheck == nil {
		errorCheck = ShouldNotHaveError()
	} else {
		errorCheck = c.ErrorCheck
	}

	t.Run(c.Name, func(t *testing.T) {
		t.Helper()
		CheckCommand(t, c.CmdBuilder(), outputCheck, errorCheck, c.Args...)
	})
}

// Run runs Run method for the all testCases.
func (cs TestCases) Run(t *testing.T) {
	t.Helper()

	for _, c := range ([]TestCase)(cs) {
		c.Run(t)
	}
}

// Build builds a testCase from the testCaseTemplate.
func (ct TestCaseTemplate) Build(cmdBuilder func() *cobra.Command) TestCase {
	return TestCase{
		Name:        ct.Name,
		CmdBuilder:  cmdBuilder,
		Args:        ct.Args,
		OutputCheck: ct.OutputCheck,
		ErrorCheck:  ct.ErrorCheck,
	}
}

// Build returns testCases.
func (cts TestCaseTemplates) Build(cmdBuilder func() *cobra.Command) TestCases {
	testCases := []TestCase{}
	for _, ct := range ([]TestCaseTemplate)(cts) {
		testCases = append(testCases, ct.Build(cmdBuilder))
	}
	return TestCases(testCases)
}

// RootSucceeds returns a template which tests if the root succeeds.
func RootSucceeds() TestCaseTemplate {
	return TestCaseTemplate{
		Name: "root succeeds",
	}
}

// HasSubcommand returns a template which tests if it has the subcommand and it succeeds.
func HasSubcommand(name string) TestCaseTemplate {
	return TestCaseTemplate{
		Name: fmt.Sprintf("has Subcommand %s and it succeeds", name),
		Args: []string{name},
	}
}
