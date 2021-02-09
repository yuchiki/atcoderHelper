package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const defaultSampleCasesDir = "sampleCases"

type opts struct {
	SampleCasesDir string
}

// Option is a functional option for NewTestCmd.
type Option func(*opts)

// SetSampleCasesDir changes sampleCaseDir from the default.
func SetSampleCasesDir(dirName string) Option {
	return func(opts *opts) {
		opts.SampleCasesDir = defaultSampleCasesDir
	}
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
		Long: `tests sample cases.

The specification is not fixed. The below is the current temporal behaviour.


- build.sh is run once/
- when there exists case{n}.input for n in 1..N, tests are done for 1..N.
- in each test,
  - the command executes "cat case{n}.input | ./run.sh > case{n}.actual".
  - Then, it compares case{n}.actual and case{n}.expected.
  - If case{n}.input is "[skip ach test]\n", the case is skipped.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("building...")
			if out, err := exec.Command("bash", "-c", "./build.sh").Output(); err != nil {
				fmt.Print(string(out))

				return err
			}
			fmt.Println("built.")

			i := 1
			successes, cases := 0, 0
			for {
				if _, err := os.Stat(testInputName(opts.SampleCasesDir, i)); err == nil {
					result, err := testNthCase(opts.SampleCasesDir, i)
					cases++
					if err != nil {
						return err
					}
					if result {
						successes++
					}
				} else {
					break
				}
				i++
			}

			fmt.Printf("total: %d/%d ", successes, cases)
			if successes == cases {
				successText("success")
			} else {
				errorText("fail")
			}

			return nil
		},
	}
}

func testInputName(sampleCasesDir string, n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.input", n))
}

func testExpectedName(sampleCasesDir string, n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.expected", n))
}

func testActualName(sampleCasesDir string, n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.actual", n))
}

func testNthCase(sampleCasesDir string, n int) (bool, error) {
	fmt.Printf("case %d: ", n)

	inputBytes, err := ioutil.ReadFile(testInputName(sampleCasesDir, n))
	if err != nil {
		return false, err
	}

	if string(inputBytes) == "[skip ach test]\n" {
		cautionText("skip")

		return true, nil
	}

	shell := fmt.Sprintf("cat %s | ./run.sh > %s", testInputName(sampleCasesDir, n), testActualName(sampleCasesDir, n))

	if err := exec.Command("bash", "-c", shell).Run(); err != nil {
		return false, err
	}

	actual, err := ioutil.ReadFile(testActualName(sampleCasesDir, n))
	if err != nil {
		return false, err
	}

	expected, err := ioutil.ReadFile(testExpectedName(sampleCasesDir, n))
	if err != nil {
		return false, err
	}

	if string(actual) == string(expected) {
		successText("pass")

		return true, nil
	}

	errorText("fail")

	if string(expected) == "" {
		fmt.Printf("  expected: (empty)\n")
	} else {
		fmt.Printf("  expected: %s", string(expected))
	}

	if string(actual) == "" {
		fmt.Printf("  actual  : (empty)\n")
	} else {
		fmt.Printf("  actual  : %s", string(actual))
	}

	return false, nil
}

var (
	errorText   = color.Red
	cautionText = color.Yellow
	successText = color.Green
)
