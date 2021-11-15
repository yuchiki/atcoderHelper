package repository

import "fmt"

const (
	AtCoderURL   = "https://atcoder.jp"
	IncomingPath = "/contests"
	RecentPath   = "/contests"
)

func taskURL(contest string, task string) string {
	return fmt.Sprintf(
		"%s/contests/%s/tasks/%s_%s",
		AtCoderURL, contest, contest, task)
}
