package testcase

type Status int

const (
	Untested Status = iota
	Pass
	NotPassed
)

func (s Status) String() string {
	switch s {
	case Untested:
		return "untested"
	case Pass:
		return "pass"
	case NotPassed:
		return "not passed"
	default:
		return "unknown status"
	}
}
