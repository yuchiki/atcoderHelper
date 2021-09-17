package testcase

type Status int

const (
	Pass Status = iota
	NotPassed
)

func (s Status) String() string {
	switch s {
	case Pass:
		return "pass"
	case NotPassed:
		return "not passed"
	default:
		return "unknown status"
	}
}
