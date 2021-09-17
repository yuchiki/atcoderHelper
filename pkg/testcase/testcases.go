package testcase

type Testcases struct {
	Testcases []Testcase
	Summary   Summary
}

type Summary struct {
	Status Status
}

func NewTestCases(bareTestcases []Testcase) Testcases {
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

	return Testcases{
		Testcases: bareTestcases,
		Summary: Summary{
			Status: status,
		},
	}
}

func (ts Testcases) MergeWithFetched(fetched Testcases) Testcases {
	// fetched の validationは省く

	unfetchedTestcases := []Testcase{}

	for _, testcase := range ts.Testcases {
		if !testcase.Fetched {
			unfetchedTestcases = append(unfetchedTestcases, testcase)
		}
	}

	joinedTestcases := append(fetched.Testcases, unfetchedTestcases...)

	return NewTestCases(joinedTestcases)
}
