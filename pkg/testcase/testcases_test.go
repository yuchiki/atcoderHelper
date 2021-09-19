package testcase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTestcases(t *testing.T) {
	testcases := []struct {
		name          string
		bareTestcases []Testcase
		testcases     Testcases
	}{
		{
			name: "all cases are passing",
			bareTestcases: []Testcase{
				{Status: Pass},
				{Status: Pass},
				{Status: Pass},
			},
			testcases: Testcases{
				Testcases: []Testcase{
					{Status: Pass},
					{Status: Pass},
					{Status: Pass},
				},
				Summary: Summary{Status: Pass},
			},
		},
		{
			name: "there exists a case not passing",
			bareTestcases: []Testcase{
				{Status: Pass},
				{Status: Pass},
				{Status: NotPassed},
			},
			testcases: Testcases{
				Testcases: []Testcase{
					{Status: Pass},
					{Status: Pass},
					{Status: NotPassed},
				},
				Summary: Summary{Status: NotPassed},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			taskTestcases := NewTestcases(testcase.bareTestcases)

			if diff := cmp.Diff(testcase.testcases, taskTestcases); diff != "" {
				t.Error(diff)
			}
		})
	}
}
