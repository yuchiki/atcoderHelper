package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const sampleCasesDir = "sampleCases"

// NewTestCmd returns test command
func NewTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "tests sample cases",
		Long:  "tests sample cases",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("building...")
			if err := exec.Command("bash", "-c", "./build.sh").Run(); err != nil {
				return err
			}
			fmt.Println("built.")

			i := 1

			successes := 0
			cases := 0
			for {
				if _, err := os.Stat(testInputName(i)); err == nil {
					result, err := testNthCase(i)
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

func testInputName(n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.input", n))
}

func testExpectedName(n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.expected", n))
}

func testActualName(n int) string {
	return path.Join(sampleCasesDir, fmt.Sprintf("case%d.actual", n))
}

func testNthCase(n int) (bool, error) {
	fmt.Printf("case %d: ", n)

	inputBytes, err := ioutil.ReadFile(testInputName(n))
	if err != nil {
		return false, err
	}

	if strings.TrimRight(string(inputBytes), "\n") == "[skip ach test]" {
		cautionText("skip")
		return true, nil
	}

	shell := fmt.Sprintf("cat %s | ./run.sh > %s", testInputName(n), testActualName(n))

	if err := exec.Command("bash", "-c", shell).Run(); err != nil {
		return false, err
	}

	actual, err := ioutil.ReadFile(testActualName(n))
	if err != nil {
		return false, err
	}

	expected, err := ioutil.ReadFile(testExpectedName(n))
	if err != nil {
		return false, err
	}

	if string(actual) == string(expected) {
		successText("pass")
		return true, nil
	} else {
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
}

var errorText = color.Red
var cautionText = color.Yellow
var successText = color.Green
