package testcase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTestcase_Check(t *testing.T) {
	testcases := []struct {
		name     string
		expected string
		actual   string
		status   Status
	}{
		{
			name:     "Pass when no diff",
			expected: "foo",
			actual:   "foo",
			status:   Pass,
		},
		{
			name:     "NotPassed with diff",
			expected: "foo",
			actual:   "bar",
			status:   NotPassed,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()

			taskTestcase := Testcase{
				Expected: testcase.expected,
			}

			_ = taskTestcase.Check(testcase.actual)

			if diff := cmp.Diff(testcase.actual, taskTestcase.Actual); diff != "" {
				t.Error(diff)
			}
			if taskTestcase.Status != testcase.status {
				t.Errorf("expected %v, but actual %v", testcase.status, taskTestcase.Status)
			}
		})
	}
}
