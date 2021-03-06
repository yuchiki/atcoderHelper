package queryablehtml

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var errMockFetchParseError = fmt.Errorf("mock fetch parse error")

func fetch(url string) (QueryableNode, error) {
	res, err := http.Get(url) //nolint:gosec
	if err != nil {
		return QueryableNode{}, fmt.Errorf("cannot open the page: %w", err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return QueryableNode{}, err
	}

	return NewQueryableNode(doc), nil
}

// SetMockFetcher replaces function Fetch to a mock fetcher function,
// that returns the result of parsing given raw HTML.
// the return value is a function that restores the value of field Fetch.
// it must be called at the last of each testcase, like that `defer callback()`
func SetMockFetcher(t *testing.T, expectedURL string, rawHTML string) func() {
	t.Helper()

	orig := Fetch

	Fetch = func(url string) (QueryableNode, error) {
		t.Helper()

		if url != expectedURL {
			t.Errorf("expected URL %s, but actual %s", expectedURL, url)
		}

		doc, err := html.Parse(strings.NewReader(rawHTML))
		if err != nil {
			t.Fatalf("fail to parse html: %s", rawHTML)
		}

		return NewQueryableNode(doc), nil
	}

	return func() { Fetch = orig }
}

// SetFailingMockFetcher replaces function Fetch to a mock fetcher function that always fails.
// the return value is a function that restores the value of field Fetch.
// it must be called at the last of each testcase, like that `defer callback()`
func SetFailingMockFetcher(t *testing.T, expectedURL string) func() {
	t.Helper()

	orig := Fetch

	Fetch = func(url string) (QueryableNode, error) {
		if url != expectedURL {
			t.Errorf("expected URL %s, but acutal %v", expectedURL, url)
		}

		return QueryableNode{}, errMockFetchParseError
	}

	return func() { Fetch = orig }
}

// Fetch is a function to fetch contents. It's return value is a queryable node.
var Fetch = fetch
