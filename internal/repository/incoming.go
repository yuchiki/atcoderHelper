package repository

import (
	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

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

	return getContests(node, "contest-table-upcoming")
}
