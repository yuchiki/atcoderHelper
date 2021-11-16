package repository

import (
	"fmt"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

// FetchTasks fetches the list of tasks in the given contest.
func FetchTasks(contest string) ([]string, error) {
	path := fmt.Sprintf("%s/contests/%s/tasks", AtCoderURL, contest)

	node, err := queryablehtml.Fetch(path)
	if err != nil {
		return nil, err
	}

	taskTRs := node.GetNodesByTag("tr")

	taskNames := []string{}
	for _, taskTR := range taskTRs {
		tds, err := taskTR.GetChildrenByTag("td")

		if err != nil {
			return nil, err
		}

		if len(tds) == 0 { // table header
			continue
		}

		taskName, err := tds[0].GetChildByTag("a").GetText()
		if err != nil {
			return nil, err
		}

		taskNames = append(taskNames, taskName)
	}

	return taskNames, nil

}
