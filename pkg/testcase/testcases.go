package testcase

import (
	"io/ioutil"
	"os"

	"github.com/yuchiki/atcoderHelper/internal/repository"
	"gopkg.in/yaml.v2"
)

var TestcasesFile = "testcases.yaml"

type Testcases struct {
	Testcases []Testcase
	Summary   Summary
}

type Summary struct {
	Status Status
	Total  int
	Passed int
}

func NewTestcases(bareTestcases []Testcase) Testcases {
	allPass := func() bool {
		for _, testcase := range bareTestcases {
			if testcase.Status != Pass {
				return false
			}
		}

		return true
	}

	var status Status
	if allPass() {
		status = Pass
	} else {
		status = NotPassed
	}

	numPassed := 0

	for _, tcase := range bareTestcases {
		if tcase.Status == Pass {
			numPassed++
		}
	}

	return Testcases{
		Testcases: bareTestcases,
		Summary: Summary{
			Status: status,
			Total:  len(bareTestcases),
			Passed: numPassed,
		},
	}
}

func (ts Testcases) MergeWithFetched(fetched []repository.Testcase) Testcases {
	// fetched の validationは省く
	unfetchedTestcases := []Testcase{}

	for _, testcase := range ts.Testcases {
		if !testcase.Fetched {
			unfetchedTestcases = append(unfetchedTestcases, testcase)
		}
	}

	fetchedTestcases := []Testcase{}

	for _, rawTestcase := range fetched {
		tcase := Testcase{
			Fetched:  true,
			Input:    rawTestcase.Input,
			Expected: rawTestcase.Expected,
			Status:   Untested,
		}

		fetchedTestcases = append(fetchedTestcases, tcase)
	}

	joinedTestcases := append(fetchedTestcases, unfetchedTestcases...) //nolint:gocritic // this is intended.

	return NewTestcases(joinedTestcases)
}

func ReadFrom(file string) (Testcases, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return Testcases{}, nil
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return Testcases{}, err
	}

	v := Testcases{}
	if err := yaml.Unmarshal(b, &v); err != nil {
		return Testcases{}, err
	}

	return v, nil
}

func (ts *Testcases) WriteTo(file string) error {
	b, err := yaml.Marshal(ts)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
