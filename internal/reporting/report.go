package reporting

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/inspection"
)

type Summary struct {
	Items   []domain.InspectionRecord `json:"items"`
	Metrics inspection.Metrics        `json:"metrics"`
}

func Build(rs []domain.InspectionRecord) Summary {
	return Summary{Items: rs, Metrics: inspection.Summarize(rs)}
}
func Filter(rs []domain.InspectionRecord, status string) []domain.InspectionRecord {
	out := []domain.InspectionRecord{}
	for _, r := range rs {
		if status == "" || r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func Titles() map[string]string {
	return map[string]string{"pending": "待检", "passed": "通过", "failed": "失败"}
}
