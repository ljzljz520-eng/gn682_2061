package reporting

import (
	"fmt"
	"inspectionbase/internal/domain"
	"inspectionbase/internal/inspection"
)

func StatusText(v string) string {
	switch v {
	case "pending":
		return "待检"
	case "passed":
		return "通过"
	case "failed":
		return "失败"
	default:
		return "未知"
	}
}
func RecordLine(r domain.InspectionRecord) string {
	return fmt.Sprintf("%s,%s,%s,%s", r.ID, r.DeviceID, StatusText(r.Status), r.Inspector)
}
func SortByID(rs []domain.InspectionRecord) []domain.InspectionRecord {
	out := append([]domain.InspectionRecord{}, rs...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].ID < out[i].ID {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
func CountPassed(rs []domain.InspectionRecord) int {
	n := 0
	for _, r := range rs {
		if r.Status == "passed" {
			n++
		}
	}
	return n
}
func CountFailed(rs []domain.InspectionRecord) int {
	n := 0
	for _, r := range rs {
		if r.Status == "failed" {
			n++
		}
	}
	return n
}
func CountPending(rs []domain.InspectionRecord) int {
	n := 0
	for _, r := range rs {
		if r.Status == "pending" {
			n++
		}
	}
	return n
}
func StatusRatio(rs []domain.InspectionRecord, status string) float64 {
	if len(rs) == 0 {
		return 0
	}
	n := 0
	for _, r := range rs {
		if r.Status == status {
			n++
		}
	}
	return float64(n) / float64(len(rs))
}
func HasFailures(rs []domain.InspectionRecord) bool { return CountFailed(rs) > 0 }
func HasPending(rs []domain.InspectionRecord) bool  { return CountPending(rs) > 0 }
func DeviceSet(rs []domain.InspectionRecord) map[string]bool {
	m := map[string]bool{}
	for _, r := range rs {
		m[r.DeviceID] = true
	}
	return m
}
func EmptySummary() Summary {
	return Summary{Items: []domain.InspectionRecord{}, Metrics: inspection.Metrics{}}
}
