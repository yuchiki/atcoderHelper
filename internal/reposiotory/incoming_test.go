package repository

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yuchiki/atcoderHelper/internal/testutil"
	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

func TestFetchIncoming(t *testing.T) {
	type testcase struct {
		name string
		// nolint:structcheck // なぜかunused判定されてしまうので
		shouldFetchParsePhaseFail bool
		html                      string
		output                    []ContestInfo
		errorCheck                testutil.ErrorCheck
	}

	testcases := []testcase{
		{
			name: "OK",
			html: `<div id="contest-table-upcoming">
					<div><div><table><tbody>
									<tr><td></td><td><a href="url1">link1</a></td></tr>
									<tr><td></td><td><a href="url2">link2</a></td></tr>
									<tr><td></td><td><a href="url3">link3</a></td></tr>
					</tbody></table></div></div>`,
			output: []ContestInfo{
				{URL: "url1", Name: "link1"},
				{URL: "url2", Name: "link2"},
				{URL: "url3", Name: "link3"},
			},
		},
		{
			name:                      "return error when failed before parsing",
			shouldFetchParsePhaseFail: true,
			errorCheck:                testutil.AnyError(),
		},
		{
			name:       "return error when parsed",
			errorCheck: testutil.AnyError(),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			callback := queryablehtml.SetMockFetcher(t, AtCoderURL+IncomingPath, testcase.html)
			defer callback()

			contestInfos, err := FetchIncoming()
			if diff := cmp.Diff(testcase.output, contestInfos); diff != "" {
				t.Error(diff)
			}

			if testcase.errorCheck == nil {
				testutil.ShouldNotHaveError()(t, err)
			} else {
				testcase.errorCheck(t, err)
			}
		})
	}
}
