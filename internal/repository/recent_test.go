package repository

import (
	"testing"
)

func TestFetchRecent(t *testing.T) {
	fetchListOfContestTest(t, FetchRecent, AtCoderURL+RecentPath, "contest-table-recent")
}
