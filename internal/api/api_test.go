package api

import (
	"inspectionbase/internal/inspection"
	"inspectionbase/internal/store"
	"net/http/httptest"
	"testing"
)

func TestHTTPRoutes(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	h := New(inspection.New(s, nil)).Routes()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
