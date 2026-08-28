package domain

func DeviceLabel(d Device) string {
	if d.Location == "" {
		return d.Name
	}
	return d.Name + " @ " + d.Location
}
func RecordLabel(r InspectionRecord) string                     { return r.ID + " / " + r.DeviceID }
func AuditLabel(a AuditEntry) string                            { return a.Action + " / " + a.RecordID }
func SettingLabel(s InspectionSetting) string                   { return s.Key + "=" + s.Value }
func IsRecordEmpty(r InspectionRecord) bool                     { return r.ID == "" && r.DeviceID == "" }
func IsDeviceEmpty(d Device) bool                               { return d.ID == "" && d.Name == "" }
func IsAuditEmpty(a AuditEntry) bool                            { return a.ID == "" && a.RecordID == "" }
func IsSettingEmpty(s InspectionSetting) bool                   { return s.Key == "" }
func RecordStatus(r InspectionRecord) Status                    { return Status(r.Status) }
func RecordIsTerminal(r InspectionRecord) bool                  { return IsTerminalStatus(r.Status) }
func RecordIsPending(r InspectionRecord) bool                   { return IsPendingStatus(r.Status) }
func DeviceCanInspect(d Device) bool                            { return CanInspect(d) }
func RecordNeedsNotes(r InspectionRecord) bool                  { return RequiresNotes(r.Status) && r.Notes == "" }
func RecordHasInspector(r InspectionRecord) bool                { return r.Inspector != "" }
func RecordHasDevice(r InspectionRecord) bool                   { return r.DeviceID != "" }
func RecordHasTimestamp(r InspectionRecord) bool                { return r.CheckedAt != "" }
func DeviceHasLocation(d Device) bool                           { return d.Location != "" }
func DeviceIsNamed(d Device) bool                               { return d.Name != "" }
func AuditHasActor(a AuditEntry) bool                           { return a.Actor != "" }
func AuditHasAction(a AuditEntry) bool                          { return a.Action != "" }
func SettingHasValue(s InspectionSetting) bool                  { return s.Value != "" }
func RecordMatchesDevice(r InspectionRecord, id string) bool    { return r.DeviceID == id }
func RecordMatchesInspector(r InspectionRecord, id string) bool { return r.Inspector == id }
func DeviceMatchesLocation(d Device, v string) bool             { return d.Location == v }
func AuditMatchesActor(a AuditEntry, v string) bool             { return a.Actor == v }
func SettingMatchesKey(s InspectionSetting, v string) bool      { return s.Key == v }
func DeviceActiveFlag(d Device) int {
	if d.Active {
		return 1
	}
	return 0
}
func RecordStatusCode(r InspectionRecord) int {
	switch r.Status {
	case "passed":
		return 2
	case "failed":
		return 3
	default:
		return 1
	}
}
func AuditPriority(a AuditEntry) int {
	if a.Action == "status" {
		return 2
	}
	return 1
}
func SettingPresent(s InspectionSetting) bool { return s.Key != "" && s.Value != "" }
func RecordSummaryParts(r InspectionRecord) []string {
	return []string{r.ID, r.DeviceID, r.Status, r.Inspector}
}
func DeviceSummaryParts(d Device) []string             { return []string{d.ID, d.Name, d.Location} }
func AuditSummaryParts(a AuditEntry) []string          { return []string{a.ID, a.Action, a.RecordID} }
func SettingSummaryParts(s InspectionSetting) []string { return []string{s.Key, s.Value, s.UpdatedAt} }
func IsSameDevice(a, b Device) bool                    { return a.ID == b.ID }
func IsSameRecord(a, b InspectionRecord) bool          { return a.ID == b.ID }
