package queryablehtml

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func fetch(url string) (QueryableNode, error) {
	res, err := http.Get(url)
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

func SetMockFetcher(expectedURL string, rawHtml string, t *testing.T) func() {
	t.Helper()

	orig := Fetch

	Fetch = func(url string) (QueryableNode, error) {
		t.Helper()

		if url != expectedURL {
			t.Errorf("expected URL %s, but actual %s", expectedURL, url)
		}

		doc, err := html.Parse(strings.NewReader(rawHtml))
		if err != nil {
			t.Fatalf("fail to parse html: %s", rawHtml)
		}

		return NewQueryableNode(doc), nil

	}

	return func() { Fetch = orig }
}

var (
	Fetch = fetch
)
