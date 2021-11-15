package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
	"github.com/yuchiki/atcoderHelper/internal/config"
	"github.com/yuchiki/atcoderHelper/pkg/testcase"
)

var (
	errorText = color.Red
	// cautionText = color.Yellow
	successText = color.Green
)

// Option is a functional option for NewTestCmd.
type Option func(*opts)

type opts struct {
	testcaseFile string
}

// NewTestCmd returns test command
func NewTestCmd(options ...Option) *cobra.Command {
	opts := opts{
		testcaseFile: testcase.TestcasesFile,
	}

	for _, option := range options {
		option(&opts)
	}

	return &cobra.Command{
		Use:   "test",
		Short: "tests sample cases",
		Long:  `tests sample cases.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			language, err := config.GetLanguage()
			if err != nil {
				return err
			}

			if err := build(language.Build); err != nil {
				return err
			}

			testcases, err := testcase.ReadFrom(opts.testcaseFile)
			if err != nil {
				return err
			}

			updatedTestcases, err := testAll(language.Run, testcases.Testcases)
			if err != nil {
				return err
			}

			showSummary(updatedTestcases)

			err = updatedTestcases.WriteTo(opts.testcaseFile)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func build(buildCommand string) error {
	fmt.Println("building...")

	if out, err := exec.Command("bash", "-c", buildCommand).Output(); err != nil {
		fmt.Print(string(out))

		return err
	}

	fmt.Println("built.")

	return nil
}

func testAll(runCommand string, testcases []testcase.Testcase) (testcase.Testcases, error) {
	updatedTestCases := []testcase.Testcase{}

	for _, singleTestcase := range testcases {
		testedTestcase, err := testSingleTestcase(runCommand, singleTestcase)
		if err != nil {
			return testcase.Testcases{}, err
		}

		updatedTestCases = append(updatedTestCases, testedTestcase)
	}

	return testcase.NewTestcases(updatedTestCases), nil
}

func testSingleTestcase(runCommand string, tcase testcase.Testcase) (testcase.Testcase, error) {
	fmt.Printf("\ntesting... ")

	updatedTestcase := testcase.Testcase{
		Fetched:  tcase.Fetched,
		Input:    tcase.Input,
		Expected: tcase.Expected,
	}

	tmpfile, err := ioutil.TempFile("", "tempfile-for-testing-")
	if err != nil {
		return testcase.Testcase{}, err
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(tcase.Input); err != nil {
		return testcase.Testcase{}, err
	}

	shell := fmt.Sprintf(
		"cat %s | %s",
		tmpfile.Name(),
		runCommand,
	)

	out, err := exec.Command("bash", "-c", shell).Output()
	if err != nil {
		updatedTestcase.Status = testcase.NotPassed

		return updatedTestcase, nil
	}

	updatedTestcase.Actual = string(out)

	if cmp.Diff(string(out), tcase.Expected) == "" {
		successText("pass")

		updatedTestcase.Status = testcase.Pass
	} else {
		errorText("fail")

		fmt.Print("  expected:\n")
		fmt.Print(indent(2, tcase.Expected)) //nolint:gomnd

		fmt.Print("  but actual:\n")
		fmt.Print(indent(2, string(out))) //nolint:gomnd

		updatedTestcase.Status = testcase.NotPassed
	}

	return updatedTestcase, nil
}

func showSummary(testcases testcase.Testcases) {
	fmt.Println()
	fmt.Println("summary:")
	fmt.Printf("%d/%d passed\n", testcases.Summary.Passed, testcases.Summary.Total)

	fmt.Print("status: ")

	if testcases.Summary.Status == testcase.Pass {
		successText(testcases.Summary.Status.String())
	} else {
		errorText(testcases.Summary.Status.String())
	}
}

func indent(indent int, text string) string {
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += "  "
	}

	hasNewLineInEnd := text[len(text)-1] == '\n'

	if hasNewLineInEnd {
		text = text[0 : len(text)-1]
	}

	lines := strings.Split(text, "\n")

	indentedLines := []string{}
	for _, line := range lines {
		indentedLines = append(indentedLines, indentation+line)
	}

	if hasNewLineInEnd {
		indentedLines = append(indentedLines, "")
	}

	return strings.Join(indentedLines, "\n")
}
