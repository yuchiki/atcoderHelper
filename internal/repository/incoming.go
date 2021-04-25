package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

const (
	AtCoderURL   = "https://atcoder.jp"
	IncomingPath = "/contests"
)

var errContestPathCannotBeParsed = errors.New("contest path cannot be parsed")

type ContestInfo struct {
	ID   string
	Name string
}

// FetchIncoming fetches information of incoming contests.
func FetchIncoming() ([]ContestInfo, error) {
	node, err := queryablehtml.Fetch(AtCoderURL + IncomingPath)
	if err != nil {
		return nil, err
	}

	contestTRs, err := node.
		GetNodeByID("contest-table-upcoming").
		GetChildByTag("div").
		GetChildByTag("div").
		GetChildByTag("table").
		GetChildByTag("tbody").
		GetChildrenByTag("tr")
	if err != nil {
		return nil, err
	}

	contestInfos := []ContestInfo{}

	for _, tr := range contestTRs {
		contestInfo, err := trToContestInfo(tr)
		if err != nil {
			return nil, err
		}

		contestInfos = append(contestInfos, contestInfo)
	}

	return contestInfos, nil
}

func trToContestInfo(tr queryablehtml.QueryableNode) (ContestInfo, error) {
	tds, err := tr.GetChildrenByTag("td")
	if err != nil {
		return ContestInfo{}, err
	}

	//nolint: gomnd
	if len(tds) < 2 {
		return ContestInfo{}, fmt.Errorf("second td does not exist: %w", queryablehtml.ErrNodeNotFound)
	}

	link := tds[1].GetChildByTag("a")
	if link.Err != nil {
		return ContestInfo{}, err
	}

	url, err := link.GetAttr("href")
	if err != nil {
		return ContestInfo{}, err
	}

	id, err := getContestID(url)
	if err != nil {
		return ContestInfo{}, err
	}

	name, err := link.GetText()
	if err != nil {
		return ContestInfo{}, err
	}

	return ContestInfo{
		ID:   id,
		Name: name,
	}, nil
}

func getContestID(contestRelativePath string) (string, error) {
	each := strings.Split(contestRelativePath, "/")

	//nolint:gomnd
	if len(each) != 3 {
		return "", fmt.Errorf(
			"path '%s' cannot be parsed: %w",
			contestRelativePath,
			errContestPathCannotBeParsed)
	}

	return each[2], nil
}
