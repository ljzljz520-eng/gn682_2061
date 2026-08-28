package inspection

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/store"
	"os"
	"testing"
)

func testService(t *testing.T) *Service {
	p := t.TempDir() + "/x.db"
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close(); os.Remove(p) })
	if e = s.SaveDevice(domain.Device{ID: "d1", Name: "泵", Active: true}); e != nil {
		t.Fatal(e)
	}
	return New(s, nil)
}
func TestInspectionNotFoundError(t *testing.T) {
	s := testService(t)
	_, e := s.Get("missing")
	if e != domain.ErrNotFound {
		t.Fatalf("want not found, got %v", e)
	}
}
func TestCreateRecord(t *testing.T) {
	s := testService(t)
	e := s.Create(domain.InspectionRecord{ID: "r1", DeviceID: "d1", Status: "pending", Inspector: "u", CheckedAt: "2025-01-01"})
	if e != nil {
		t.Fatal(e)
	}
}
