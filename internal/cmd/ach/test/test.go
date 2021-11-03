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
	"gopkg.in/yaml.v2"
)

const (
	TestcasesFile = "testcases.yaml"
)

var (
	errorText   = color.Red
	cautionText = color.Yellow
	successText = color.Green
)

// Option is a functional option for NewTestCmd.
type Option func(*opts)

type opts struct {
	SampleCasesDir string
}

type summary struct {
	total int
	pass  int
}

// NewTestCmd returns test command
func NewTestCmd(options ...Option) *cobra.Command {
	opts := opts{
		SampleCasesDir: "sampleCases",
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

			testcases, err := readTestCases(TestcasesFile)
			if err != nil {
				return err
			}

			updatedTestcases, err := testAll(language.Run, testcases.Testcases)
			if err != nil {
				return err
			}

			showSummary(updatedTestcases)

			err = writeTestcases(updatedTestcases, TestcasesFile)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func build(buildCommand string) error {
	fmt.Println("building...")

	if out, err := exec.Command("bash", "-c", buildCommand).Output(); err != nil { //nolint: gosec
		fmt.Print(string(out))

		return err
	}

	fmt.Println("built.")

	return nil
}

func readTestCases(file string) (testcase.Testcases, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return testcase.Testcases{}, err
	}

	v := testcase.Testcases{}
	if err := yaml.Unmarshal(b, &v); err != nil {
		return testcase.Testcases{}, err
	}

	return v, nil
}

func writeTestcases(testcases testcase.Testcases, testcasesFile string) error {
	b, err := yaml.Marshal(testcases)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(testcasesFile, b, os.ModeExclusive)
	if err != nil {
		return err
	}

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

	tmpfile.WriteString(tcase.Input)

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
		fmt.Print(indent(2, tcase.Expected))

		fmt.Print("  but actual:\n")
		fmt.Print(indent(2, string(out)))

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
		indentation = indentation + "  "
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
