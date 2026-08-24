package inspection

import "inspectionbase/internal/domain"

type Metrics struct{ Total, Passed, Failed, Pending int }

func Summarize(rs []domain.InspectionRecord) Metrics {
	m := Metrics{Total: len(rs)}
	for _, r := range rs {
		switch r.Status {
		case "passed":
			m.Passed++
		case "failed":
			m.Failed++
		case "pending":
			m.Pending++
		}
	}
	return m
}
func Completion(m Metrics) float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Passed+m.Failed) / float64(m.Total)
}
