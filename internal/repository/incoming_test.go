package repository

import (
	"testing"
)

func TestFetchIncoming(t *testing.T) {
	fetchListOfContestTest(t, FetchIncoming, AtCoderURL+IncomingPath, "contest-table-upcoming")
}
