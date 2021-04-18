package queryablehtml

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func parseHTML(expression string, t *testing.T) QueryableNode {
	t.Helper()

	doc, err := html.Parse(strings.NewReader(expression))
	if err != nil {
		t.Fatal(err)
	}

	return NewQueryableNode(doc)
}

func inBody(expression string) string {
	return "<html><head></head><body>" + expression + "</body></html>"
}

func TestGetNodeByID(t *testing.T) {
	t.Fatal("not implemented yet")
}

func TestGetChildrenByTag(t *testing.T) {
	t.Fatal("not implemented yet")
}

func TestGetChildByTag(t *testing.T) {
	t.Fatal("not implemented yet")
}

func TestGetAttr(t *testing.T) {
	t.Fatal("not implemented yet")
}

func TestGetText(t *testing.T) {
	query := func(node QueryableNode) (string, error) {
		return node.GetNodeByID("root").GetText()
	}

	type testcase struct {
		name     string
		html     string
		expected string
		err      error
	}

	testcases := []testcase{
		{
			name:     "OK",
			html:     "<div id=\"root\">foo</div>",
			expected: "foo",
		},
		{
			name: "error when the node is not a text node",
			html: "<div id=\"root\"><div></div></div>",
			err:  ErrNotTextNode,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(testcase.html, t)
			expected := testcase.expected
			actual, err := query(node)
			if !errors.Is(err, testcase.err) {
				t.Fatalf("expected err %v, but actual err %v", testcase.err, err)
			}

			if actual != expected {
				t.Errorf("expected %s, but actual %s", expected, actual)
			}
		})
	}

}

func TestString(t *testing.T) {
	nodeString := "<div>foo</div>"
	expected := inBody(nodeString)

	node := parseHTML(nodeString, t)
	actual := node.String()

	if actual != expected {
		t.Errorf("expected %v, but actual %v", expected, actual)
	}
}
