package reporting

import "inspectionbase/internal/domain"

type DeviceSummary struct {
	DeviceID              string
	Total, Passed, Failed int
}

func SummarizeDevices(rs []domain.InspectionRecord) []DeviceSummary {
	m := map[string]*DeviceSummary{}
	for _, r := range rs {
		x := m[r.DeviceID]
		if x == nil {
			x = &DeviceSummary{DeviceID: r.DeviceID}
			m[r.DeviceID] = x
		}
		x.Total++
		if r.Status == "passed" {
			x.Passed++
		}
		if r.Status == "failed" {
			x.Failed++
		}
	}
	out := []DeviceSummary{}
	for _, x := range m {
		out = append(out, *x)
	}
	return out
}
func PassRate(s DeviceSummary) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Passed) / float64(s.Total)
}
func FailedDevices(v []DeviceSummary) []DeviceSummary {
	out := []DeviceSummary{}
	for _, x := range v {
		if x.Failed > 0 {
			out = append(out, x)
		}
	}
	return out
}
func HealthyDevices(v []DeviceSummary) []DeviceSummary {
	out := []DeviceSummary{}
	for _, x := range v {
		if x.Failed == 0 && x.Total > 0 {
			out = append(out, x)
		}
	}
	return out
}
func TotalChecks(v []DeviceSummary) int {
	n := 0
	for _, x := range v {
		n += x.Total
	}
	return n
}
func TotalFailures(v []DeviceSummary) int {
	n := 0
	for _, x := range v {
		n += x.Failed
	}
	return n
}
func TotalPasses(v []DeviceSummary) int {
	n := 0
	for _, x := range v {
		n += x.Passed
	}
	return n
}
func IsReliable(x DeviceSummary) bool { return x.Total >= 3 && PassRate(x) >= 0.8 }
func ReliableDevices(v []DeviceSummary) []DeviceSummary {
	out := []DeviceSummary{}
	for _, x := range v {
		if IsReliable(x) {
			out = append(out, x)
		}
	}
	return out
}
func NeedsAttention(x DeviceSummary) bool { return x.Failed > 0 || x.Total == 0 }
func AttentionDevices(v []DeviceSummary) []DeviceSummary {
	out := []DeviceSummary{}
	for _, x := range v {
		if NeedsAttention(x) {
			out = append(out, x)
		}
	}
	return out
}
func StatusCounts(rs []domain.InspectionRecord) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func Inspectors(rs []domain.InspectionRecord) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Inspector]++
	}
	return m
}
func DeviceIDs(rs []domain.InspectionRecord) []string {
	m := DeviceSet(rs)
	out := []string{}
	for id := range m {
		out = append(out, id)
	}
	return out
}
func HasDevice(rs []domain.InspectionRecord, id string) bool {
	for _, r := range rs {
		if r.DeviceID == id {
			return true
		}
	}
	return false
}
