package repository

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yuchiki/atcoderHelper/internal/testutil"
	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

type contestsFetcher func() ([]ContestInfo, error)

func fetchListOfContestTest(t *testing.T, toBeTested contestsFetcher, url string, nodeID string) {
	t.Helper()

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
			html: fmt.Sprintf(`<div id="%s">
					<div><div><table><tbody>
									<tr><td></td><td><a href="/contest/id1">link1</a></td></tr>
									<tr><td></td><td><a href="/contest/id2">link2</a></td></tr>
									<tr><td></td><td><a href="/contest/id3">link3</a></td></tr>
					</tbody></table></div></div>`, nodeID),
			output: []ContestInfo{
				{ID: "id1", Name: "link1"},
				{ID: "id2", Name: "link2"},
				{ID: "id3", Name: "link3"},
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
			callback := queryablehtml.SetMockFetcher(t, url, testcase.html)
			defer callback()

			contestInfos, err := toBeTested()
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
