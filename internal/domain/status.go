package domain

type Status string

const (
	StatusPending Status = "pending"
	StatusPassed  Status = "passed"
	StatusFailed  Status = "failed"
)

func (s Status) Terminal() bool { return s == StatusPassed || s == StatusFailed }
func ParseStatus(v string) (Status, bool) {
	s := Status(v)
	switch s {
	case StatusPending, StatusPassed, StatusFailed:
		return s, true
	default:
		return "", false
	}
}
func StatusLabel(s Status) string {
	switch s {
	case StatusPending:
		return "待检"
	case StatusPassed:
		return "通过"
	case StatusFailed:
		return "失败"
	default:
		return "未知"
	}
}
