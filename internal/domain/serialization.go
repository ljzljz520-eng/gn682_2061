package domain

import (
	"encoding/json"
	"fmt"
)

func EncodeRecord(r InspectionRecord) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (InspectionRecord, error) {
	var r InspectionRecord
	err := json.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}
	return r, r.Validate()
}
func RecordKey(r InspectionRecord) string { return fmt.Sprintf("record:%s:%s", r.DeviceID, r.ID) }
func CloneRecord(r InspectionRecord) InspectionRecord {
	return InspectionRecord{ID: r.ID, DeviceID: r.DeviceID, Status: r.Status, Inspector: r.Inspector, Notes: r.Notes, CheckedAt: r.CheckedAt}
}
