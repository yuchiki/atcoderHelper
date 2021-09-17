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
