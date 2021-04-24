package repository

import (
	"fmt"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

const (
	AtCoderURL   = "https://atcoder.jp"
	IncomingPath = "/contests"
)

type ContestInfo struct {
	URL  string
	Name string
}

// URLを渡すとQueryableNodesを返す関数を切り出し
// 何かのパッケージのデフォルトにし
// テスト時に差し替えられるようにする

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

	trToContestInfo := func(tr queryablehtml.QueryableNode) (ContestInfo, error) {
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

		name, err := link.GetText()
		if err != nil {
			return ContestInfo{}, err
		}

		return ContestInfo{
			URL:  url,
			Name: name,
		}, nil
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
