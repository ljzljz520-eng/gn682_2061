package domain

type QueryFilter struct {
	DeviceID, Status, Inspector string
	Limit                       int
}

func (f QueryFilter) Matches(r InspectionRecord) bool {
	if f.DeviceID != "" && f.DeviceID != r.DeviceID {
		return false
	}
	if f.Status != "" && f.Status != r.Status {
		return false
	}
	if f.Inspector != "" && f.Inspector != r.Inspector {
		return false
	}
	return true
}
func NormalizeLimit(v int) int {
	if v <= 0 {
		return 20
	}
	if v > 100 {
		return 100
	}
	return v
}
