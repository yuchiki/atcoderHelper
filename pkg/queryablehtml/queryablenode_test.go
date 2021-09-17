package queryablehtml

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func parseHTML(t *testing.T, expression string) QueryableNode {
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
	query := func(node QueryableNode) QueryableNode {
		return node.GetNodeByID("bar")
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
			html:     `<div><a href="dummy">foo</a><div><div id="bar">baz</div></div></div>`,
			expected: `<div id="bar">baz</div>`,
		},
		{
			name: "error when it does not include nodes with id 'bar'",
			html: `<div id="root"><a href="dummy">foo</a></div>`,
			err:  ErrNodeNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(t, testcase.html)
			expected := testcase.expected
			targetNode := query(node)
			if !errors.Is(targetNode.Err, testcase.err) {
				t.Fatalf("expected err %v, but actual err %v", testcase.err, targetNode.Err)
			}

			if targetNode.Err == nil && targetNode.String() != expected {
				t.Errorf("expected %s, but actual %s", expected, targetNode.String())
			}
		})
	}
}

func TestGetChildrenByTag(t *testing.T) {
	query := func(node QueryableNode) ([]QueryableNode, error) {
		return node.GetNodeByID("root").GetChildrenByTag("div")
	}

	type testcase struct {
		name     string
		html     string
		expected []string
		err      error
	}

	testcases := []testcase{
		{
			name: "OK",
			html: `
	<div id="root">
		<div>div1</div>
		<a href="dummy">a1</a>
		<div>div2</div>
		<div>div3</div>
	</div>
`,
			expected: []string{
				"<div>div1</div>",
				"<div>div2</div>",
				"<div>div3</div>",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(t, testcase.html)
			expected := testcase.expected
			targetNodes, err := query(node)

			var actual []string
			for _, targetNode := range targetNodes {
				actual = append(actual, targetNode.String())
			}

			if !errors.Is(err, testcase.err) {
				t.Fatalf("expected err %v, but actual err %v", testcase.err, err)
			}

			diff := cmp.Diff(actual, expected)
			if err == nil && diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestGetChildByTag(t *testing.T) {
	query := func(node QueryableNode) QueryableNode {
		return node.GetNodeByID("root").GetChildByTag("div")
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
			html:     `<div id="root"><a href="dummy">foo</a><div>bar</div></div>`,
			expected: "<div>bar</div>",
		},
		{
			name: "error when the node has not attribute 'foo'",
			html: `<div id="root"><a href="dummy">foo</a></div>`,
			err:  ErrNodeNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(t, testcase.html)
			expected := testcase.expected
			targetNode := query(node)
			if !errors.Is(targetNode.Err, testcase.err) {
				t.Fatalf("expected err %v, but actual err %v", testcase.err, targetNode.Err)
			}

			if targetNode.Err == nil && targetNode.String() != expected {
				t.Errorf("expected %s, but actual %s", expected, targetNode.String())
			}
		})
	}
}

func TestGetAttr(t *testing.T) {
	query := func(node QueryableNode) (string, error) {
		return node.GetNodeByID("root").GetAttr("foo")
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
			html:     `<div id="root" foo="bar"></div>`,
			expected: "bar",
		},
		{
			name: "error when the node has not attribute 'foo'",
			html: `<div id="root"><</div>`,
			err:  ErrAttrNotFound,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(t, testcase.html)
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

func TestGetNodesByTag(t *testing.T) {
	query := func(node QueryableNode) []string {
		texts := []string{}

		for _, pre := range node.GetNodesByTag("pre") {
			text, _ := pre.GetText()
			texts = append(texts, text)
		}

		return texts
	}

	input := `
		<div>
			<pre>foo</pre>
			<div>
				<pre>bar</pre>
				<pre>baz</pre>
			</div>
			<pre>qux</pre>
		</div>
	`

	node := parseHTML(t, input)
	expected := []string{"foo", "bar", "baz", "qux"}

	if diff := cmp.Diff(query(node), expected); diff != "" {
		t.Error(diff)
	}
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
			html:     `<div id="root">foo</div>`,
			expected: "foo",
		},
		{
			name: "error when the node is not a text node",
			html: `<div id="root"><div></div></div>`,
			err:  ErrNotTextNode,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			t.Helper()
			node := parseHTML(t, testcase.html)
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

	node := parseHTML(t, nodeString)
	actual := node.String()

	if actual != expected {
		t.Errorf("expected %v, but actual %v", expected, actual)
	}
}
