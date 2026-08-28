package domain

func IsKnownStatus(s string) bool { return s == "pending" || s == "passed" || s == "failed" }
func IsActive(d Device) bool      { return d.Active }
func CanInspect(d Device) bool    { return d.Active && d.ID != "" }
func RequiresNotes(s string) bool { return s == "failed" }
func CheckNotes(s, status string) bool {
	if RequiresNotes(status) {
		return len(s) > 0
	}
	return true
}
func CheckInspector(v string) bool  { return len(v) >= 1 && len(v) <= 80 }
func CheckID(v string) bool         { return len(v) >= 1 && len(v) <= 128 }
func CheckLocation(v string) bool   { return len(v) <= 200 }
func CheckDeviceName(v string) bool { return len(v) >= 1 && len(v) <= 120 }
func NormalizeStatus(v string) string {
	if IsKnownStatus(v) {
		return v
	}
	return "pending"
}
func NormalizeActor(v string) string {
	if v == "" {
		return "system"
	}
	return v
}
func IsTerminalStatus(v string) bool { return v == "passed" || v == "failed" }
func IsPendingStatus(v string) bool  { return v == "pending" }
func IsFailure(v string) bool        { return v == "failed" }
func IsSuccess(v string) bool        { return v == "passed" }
func ValidRecord(r InspectionRecord) bool {
	return r.Validate() == nil && CheckNotes(r.Notes, r.Status) && CheckInspector(r.Inspector)
}
func ValidDevice(d Device) bool {
	return d.Validate() == nil && CheckLocation(d.Location) && CheckDeviceName(d.Name)
}
func CompareRecords(a, b InspectionRecord) bool {
	return a.ID == b.ID && a.Status == b.Status && a.DeviceID == b.DeviceID
}
func CompareDevices(a, b Device) bool {
	return a.ID == b.ID && a.Name == b.Name && a.Active == b.Active
}
func AuditActionValid(a AuditEntry) bool { return a.ID != "" && a.Action != "" && a.RecordID != "" }
