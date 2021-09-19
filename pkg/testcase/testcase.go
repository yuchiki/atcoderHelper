package testcase

type Testcase struct {
	Fetched  bool
	Input    string
	Expected string
	Actual   string
	Status   Status
}

func (t *Testcase) Check(actual string) Status {
	var status Status
	if t.Expected == actual {
		status = Pass
	} else {
		status = NotPassed
	}

	t.Actual = actual
	t.Status = status

	return status
}
