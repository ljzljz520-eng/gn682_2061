package audit

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/store"
	"testing"
)

func TestAuditTrail(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	a := New(s)
	if e = a.Record(domain.AuditEntry{ID: "a", Action: "create", RecordID: "r"}); e != nil {
		t.Fatal(e)
	}
	v, e := a.ForRecord("r")
	if e != nil || len(v) != 1 {
		t.Fatal(e)
	}
}
