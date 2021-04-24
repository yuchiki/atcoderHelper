package queryablehtml

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

var (
	ErrAttrNotFound = errors.New("attr not found")
	ErrNodeNotFound = errors.New("node not found")
	ErrNotTextNode  = errors.New("node is not text node")
)

type QueryableNode struct {
	Node *html.Node
	Err  error
}

func NewQueryableNode(node *html.Node) QueryableNode {
	return QueryableNode{node, nil}
}

func (n QueryableNode) GetNodeByID(id string) QueryableNode {
	if n.Err != nil {
		return n
	}

	targetNode, err := getNodeByID(n.Node, id)

	return QueryableNode{targetNode, err}
}

func (n QueryableNode) GetChildrenByTag(tag string) ([]QueryableNode, error) {
	if n.Err != nil {
		return nil, n.Err
	}

	nodes := getChildrenByTag(n.Node, tag)

	queryableNodes := []QueryableNode{}
	for _, node := range nodes {
		queryableNodes = append(queryableNodes, QueryableNode{node, nil})
	}

	return queryableNodes, nil
}

func (n QueryableNode) GetChildByTag(tag string) QueryableNode {
	if n.Err != nil {
		return n
	}

	targetNode, err := getChildByTag(n.Node, tag)

	return QueryableNode{targetNode, err}
}

func (n QueryableNode) GetAttr(key string) (string, error) {
	if n.Err != nil {
		return "", n.Err
	}

	return getAttr(n.Node, key)
}

func (n QueryableNode) GetText() (string, error) {
	if n.Err != nil {
		return "", n.Err
	}

	child := n.Node.FirstChild

	if child.Type != html.TextNode {
		return "", fmt.Errorf("%v is not a text node: %w", child, ErrNotTextNode)
	}

	return child.Data, nil
}

func (n QueryableNode) String() string {
	if n.Err != nil {
		return n.Err.Error()
	}

	return nodeToString(n.Node)
}

func getNodeByID(node *html.Node, id string) (*html.Node, error) {
	if getID(node) == id {
		return node, nil
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		targetNode, err := getNodeByID(child, id)

		if errors.Is(err, ErrNodeNotFound) {
			continue
		}

		if err != nil {
			return nil, err
		}

		return targetNode, nil
	}

	return nil, fmt.Errorf("node with id '%s' is not found in children of %s: %w", id, nodeToString(node), ErrNodeNotFound)
}

func getAttr(node *html.Node, key string) (string, error) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}

	return "", fmt.Errorf("attr '%s' not found in %s: %w", key, nodeToString(node), ErrAttrNotFound)
}

func getID(node *html.Node) string {
	id, _ := getAttr(node, "id")

	return id
}

func getChildByTag(node *html.Node, tag string) (*html.Node, error) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Data == tag {
			return child, nil
		}
	}

	return nil, fmt.Errorf("node with tag '%s' is not found in children of %s: %w",
		tag,
		nodeToString(node),
		ErrNodeNotFound)
}

func getChildrenByTag(node *html.Node, tag string) []*html.Node {
	var children []*html.Node

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Data == tag {
			children = append(children, child)
		}
	}

	return children
}

func nodeToString(node *html.Node) string {
	buf := new(bytes.Buffer)

	err := html.Render(buf, node)
	if err != nil {
		return err.Error()
	}

	return buf.String()
}
