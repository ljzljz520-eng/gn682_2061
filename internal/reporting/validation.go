package reporting

import "inspectionbase/internal/domain"

func ValidItems(v []domain.InspectionRecord) bool {
	for _, r := range v {
		if !domain.ValidRecord(r) {
			return false
		}
	}
	return true
}
func InvalidItems(v []domain.InspectionRecord) []domain.InspectionRecord {
	out := []domain.InspectionRecord{}
	for _, r := range v {
		if !domain.ValidRecord(r) {
			out = append(out, r)
		}
	}
	return out
}
func RecordIDs(v []domain.InspectionRecord) []string {
	out := []string{}
	for _, r := range v {
		out = append(out, r.ID)
	}
	return out
}
func DistinctIDs(v []domain.InspectionRecord) bool {
	m := map[string]bool{}
	for _, r := range v {
		if m[r.ID] {
			return false
		}
		m[r.ID] = true
	}
	return true
}
func DeviceCoverage(v []domain.InspectionRecord, ids []string) float64 {
	if len(ids) == 0 {
		return 0
	}
	n := 0
	for _, id := range ids {
		if HasDevice(v, id) {
			n++
		}
	}
	return float64(n) / float64(len(ids))
}
func PendingOnly(v []domain.InspectionRecord) bool {
	for _, r := range v {
		if r.Status != "pending" {
			return false
		}
	}
	return true
}
func TerminalOnly(v []domain.InspectionRecord) bool {
	for _, r := range v {
		if !domain.IsTerminalStatus(r.Status) {
			return false
		}
	}
	return true
}
func AnyStatus(v []domain.InspectionRecord, s string) bool {
	for _, r := range v {
		if r.Status == s {
			return true
		}
	}
	return false
}
func AllDevice(v []domain.InspectionRecord, id string) bool {
	for _, r := range v {
		if r.DeviceID != id {
			return false
		}
	}
	return true
}
func NotesPresent(v []domain.InspectionRecord) bool {
	for _, r := range v {
		if r.Notes == "" {
			return false
		}
	}
	return true
}
func InspectorsPresent(v []domain.InspectionRecord) bool {
	for _, r := range v {
		if r.Inspector == "" {
			return false
		}
	}
	return true
}
func LatestByDevice(v []domain.InspectionRecord, id string) (domain.InspectionRecord, bool) {
	var out domain.InspectionRecord
	ok := false
	for _, r := range v {
		if r.DeviceID == id {
			out = r
			ok = true
		}
	}
	return out, ok
}
