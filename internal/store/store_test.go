package store

import (
	"inspectionbase/internal/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	_ = s.SaveDevice(domain.Device{ID: "d", Name: "设备"})
	_ = s.SaveRecord(domain.InspectionRecord{ID: "r", DeviceID: "d", Status: "passed", Inspector: "u"})
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.GetRecord("r")
	if e != nil || r.ID != "r" {
		t.Fatal(e, r)
	}
}
func TestStoreSetting(t *testing.T) {
	s, e := Open(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.SaveSetting(domain.InspectionSetting{Key: "k", Value: "v"}); e != nil {
		t.Fatal(e)
	}
}
