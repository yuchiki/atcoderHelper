package testutil

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
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

	// HelpInfo contains information obtained from help message.
	HelpInfo struct {
		CommandName string
		Subcommands []string
	}
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

// HelpShouldShowCommandName returns a check for command name.
func HelpShouldShowCommandName(name string) OutputCheck {
	return func(t *testing.T, output string) {
		t.Helper()

		help := ParseHelp(t, output)

		if help.CommandName != name {
			t.Errorf("expected %v, but actual %v", name, help.CommandName)
		}
	}
}

// HelpShouldShowSubcommands returns a check for subcommands.
func HelpShouldShowSubcommands(subcommands []string) OutputCheck {
	return func(t *testing.T, output string) {
		t.Helper()

		help := ParseHelp(t, output)

		for _, subcommand := range subcommands {
			if !Contains(subcommand, help.Subcommands) {
				t.Errorf("subcommand %v is not included in %v", subcommand, help.Subcommands)
			}
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
	return func(t *testing.T, _ string) {
		t.Helper()
	}
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

func AnyError() ErrorCheck {
	return func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Error("it must have any errors")
		}
	}
}

// CheckCommand checks if the command behaves expectedly.
func CheckCommand(
	t *testing.T,
	command *cobra.Command,
	outputCheck OutputCheck,
	errorCheck ErrorCheck,
	args ...string) {
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

// HasName returns a template which tests if the root -h succeeds and shows the name.
func HasName(name string) TestCaseTemplate {
	return TestCaseTemplate{
		Name:        fmt.Sprintf("root has Name %s", name),
		Args:        []string{"-h"},
		OutputCheck: HelpShouldShowCommandName(name),
	}
}

// HasSubcommands returns a template which tests if it has the subcommand and it succeeds.
func HasSubcommands(subcommands ...string) TestCaseTemplate {
	return TestCaseTemplate{
		Name:        fmt.Sprintf("has Subcommands %s and it succeeds", subcommands),
		Args:        []string{"-h"},
		OutputCheck: HelpShouldShowSubcommands(subcommands),
	}
}

// ParseHelp parses a help message.
func ParseHelp(t *testing.T, helpMessage string) HelpInfo {
	t.Helper()

	lines := strings.Split(helpMessage, "\n")

	literal := map[string][]string{}

	i := 0
	for i < len(lines) {
		if lines[i] == "" {
			i++

			continue
		}

		if !strings.HasSuffix(lines[i], ":") {
			i++

			continue
		}

		key := strings.TrimSuffix(lines[i], ":")
		i++

		for j := i; j < len(lines); j++ {
			if !strings.HasPrefix(lines[j], "  ") {
				break
			}

			literal[key] = append(literal[key], strings.TrimSpace(lines[j]))
		}
	}

	if len(literal["Usage"]) == 0 {
		t.Fatal("help message has no Usage section")
	}

	cmdName := strings.Fields(literal["Usage"][0])[0]

	subcommands := []string{}
	for _, cmdDescriptions := range literal["Available Commands"] {
		subcommands = append(subcommands, strings.Fields(cmdDescriptions)[0])
	}

	return HelpInfo{
		CommandName: cmdName,
		Subcommands: subcommands,
	}
}

func Contains(str string, arr []string) bool {
	for _, elem := range arr {
		if str == elem {
			return true
		}
	}

	return false
}
