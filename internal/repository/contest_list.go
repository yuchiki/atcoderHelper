package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

var errContestPathCannotBeParsed = errors.New("contest path cannot be parsed")

func getContests(node queryablehtml.QueryableNode, nodeID string) ([]ContestInfo, error) {
	contestTRs, err := node.
		GetNodeByID(nodeID).
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
