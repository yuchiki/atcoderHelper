package repository

import (
	"regexp"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

type Testcase struct {
	Input    string
	Expected string
}

var (
	inputSampleRegex    = regexp.MustCompile(`入力例 (\d+)`)
	expectedSampleRegex = regexp.MustCompile(`出力例 (\d+)`)
)

func FetchTestcases(contest string, task string) ([]Testcase, error) {
	url := taskURL(contest, task)

	node, err := queryablehtml.Fetch(url)
	if err != nil {
		return nil, err
	}

	inputs := []string{}
	for _, section := range node.GetNodesByTag("section") {
		title, err := section.GetChildByTag("h3").GetText()
		if err != nil {
			break
		}

		if inputSampleRegex.MatchString(title) {
			text, err := section.GetChildByTag("pre").GetText()
			if err != nil {
				return nil, err
			}

			inputs = append(inputs, text)
		}
	}

	expecteds := []string{}
	for _, section := range node.GetNodesByTag("section") {
		title, err := section.GetChildByTag("h3").GetText()
		if err != nil {
			break
		}

		if expectedSampleRegex.MatchString(title) {
			text, err := section.GetChildByTag("pre").GetText()
			if err != nil {
				return nil, err
			}

			expecteds = append(expecteds, text)
		}
	}

	testcases := []Testcase{}
	for i := 0; i < len(inputs); i++ {
		testcases = append(testcases, Testcase{inputs[i], expecteds[i]})
	}

	return testcases, nil
}
