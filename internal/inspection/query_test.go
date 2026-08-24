package inspection

import (
	"inspectionbase/internal/domain"
	"testing"
)

func TestSecondaryWorkflow(t *testing.T) {
	s := testService(t)
	_ = s.Create(domain.InspectionRecord{ID: "r1", DeviceID: "d1", Status: "pending", Inspector: "u"})
	r, e := s.Get("r1")
	if e != nil || r.ID != "r1" {
		t.Fatalf("%v %#v", e, r)
	}
}
func TestListFilter(t *testing.T) {
	s := testService(t)
	_ = s.Create(domain.InspectionRecord{ID: "r1", DeviceID: "d1", Status: "pending", Inspector: "u"})
	v, e := s.List(domain.QueryFilter{Status: "pending"})
	if e != nil || len(v) != 1 {
		t.Fatal(e, len(v))
	}
}
