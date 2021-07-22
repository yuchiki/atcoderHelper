package repository

import (
	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

// FetchRecent fetches information of recent contests.
func FetchRecent() ([]ContestInfo, error) {
	node, err := queryablehtml.Fetch(AtCoderURL + RecentPath)
	if err != nil {
		return nil, err
	}

	return getContests(node, "contest-table-recent")
}
