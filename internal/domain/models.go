package domain

import "errors"

var ErrNotFound = errors.New("inspection record not found")
var ErrInvalid = errors.New("invalid inspection data")

type InspectionRecord struct{ ID, DeviceID, Status, Inspector, Notes, CheckedAt string }
type Device struct {
	ID, Name, Location string
	Active             bool
}
type AuditEntry struct{ ID, Action, RecordID, Actor, CreatedAt string }
type InspectionSetting struct{ Key, Value, UpdatedAt string }

func (r InspectionRecord) Validate() error {
	if r.ID == "" || r.DeviceID == "" || r.Inspector == "" {
		return ErrInvalid
	}
	if r.Status != "pending" && r.Status != "passed" && r.Status != "failed" {
		return ErrInvalid
	}
	return nil
}
func (d Device) Validate() error {
	if d.ID == "" || d.Name == "" {
		return ErrInvalid
	}
	return nil
}
func ValidTransition(from, to string) bool {
	if from == to {
		return true
	}
	if from == "pending" && (to == "passed" || to == "failed") {
		return true
	}
	return false
}
