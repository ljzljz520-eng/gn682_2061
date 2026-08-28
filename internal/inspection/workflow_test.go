package inspection

import (
	"inspectionbase/internal/domain"
	"testing"
)

func TestPrimaryWorkflow(t *testing.T) {
	s := testService(t)
	if e := s.Create(domain.InspectionRecord{ID: "r1", DeviceID: "d1", Status: "pending", Inspector: "u"}); e != nil {
		t.Fatal(e)
	}
}
func TestTertiaryWorkflow(t *testing.T) {
	s := testService(t)
	_ = s.Create(domain.InspectionRecord{ID: "r1", DeviceID: "d1", Status: "pending", Inspector: "u"})
	if e := s.UpdateStatus("r1", "passed", "u"); e != nil {
		t.Fatal(e)
	}
}
