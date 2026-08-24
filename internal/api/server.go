package api

import (
	"encoding/json"
	"errors"
	"inspectionbase/internal/domain"
	"inspectionbase/internal/inspection"
	"inspectionbase/internal/reporting"
	"net/http"
	"strings"
)

type Server struct{ svc *inspection.Service }

func New(s *inspection.Service) *Server { return &Server{svc: s} }
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/api/inspections", s.collection)
	m.HandleFunc("/api/inspections/", s.item)
	return m
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{"status":"ok"}`))
}
func (s *Server) collection(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		items, e := s.svc.List(domain.QueryFilter{Limit: 20})
		if e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 200, reporting.Build(items))
		return
	}
	if r.Method == "POST" {
		var rec domain.InspectionRecord
		if json.NewDecoder(r.Body).Decode(&rec) != nil {
			writeErr(w, domain.ErrInvalid)
			return
		}
		if e := s.svc.Create(rec); e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 201, rec)
		return
	}
	w.WriteHeader(405)
}
func (s *Server) item(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/inspections/")
	if r.Method == "GET" {
		rec, e := s.svc.Get(id)
		if e != nil {
			writeErr(w, e)
			return
		}
		if rec.ID == "" {
			writeErr(w, domain.ErrNotFound)
			return
		}
		writeJSON(w, 200, rec)
		return
	}
	if r.Method == "PATCH" {
		var p struct{ Status, Actor string }
		if json.NewDecoder(r.Body).Decode(&p) != nil {
			writeErr(w, domain.ErrInvalid)
			return
		}
		if e := s.svc.UpdateStatus(id, p.Status, p.Actor); e != nil {
			writeErr(w, e)
			return
		}
		writeJSON(w, 200, map[string]string{"status": "updated"})
		return
	}
	w.WriteHeader(405)
}
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
func writeErr(w http.ResponseWriter, e error) {
	status := 500
	if errors.Is(e, domain.ErrNotFound) {
		status = 404
	}
	if errors.Is(e, domain.ErrInvalid) {
		status = 400
	}
	writeJSON(w, status, map[string]string{"error": e.Error()})
}
