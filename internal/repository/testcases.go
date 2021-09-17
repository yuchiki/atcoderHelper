package repository

import (
	"regexp"

	"github.com/yuchiki/atcoderHelper/pkg/queryablehtml"
)

type Testcase struct {
	Input  string
	Output string
}

var (
	inputSampleRegex  = regexp.MustCompile(`入力例 (\d+)`)
	outputSampleRegex = regexp.MustCompile(`出力例 (\d+)`)
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

	outputs := []string{}
	for _, section := range node.GetNodesByTag("section") {
		title, err := section.GetChildByTag("h3").GetText()
		if err != nil {
			break
		}

		if outputSampleRegex.MatchString(title) {
			text, err := section.GetChildByTag("pre").GetText()
			if err != nil {
				return nil, err
			}

			outputs = append(outputs, text)
		}
	}

	testcases := []Testcase{}
	for i := 0; i < len(inputs); i++ {
		testcases = append(testcases, Testcase{inputs[i], outputs[i]})
	}

	return testcases, nil
}
