package inspection

import (
	"fmt"
	"inspectionbase/internal/domain"
	"inspectionbase/internal/store"
)

type Service struct {
	st    *store.Store
	audit func(domain.AuditEntry) error
}

func New(st *store.Store, a func(domain.AuditEntry) error) *Service {
	return &Service{st: st, audit: a}
}
func (s *Service) Create(r domain.InspectionRecord) error {
	if e := r.Validate(); e != nil {
		return e
	}
	if _, e := s.st.GetDevice(r.DeviceID); e != nil {
		return e
	}
	if e := s.st.SaveRecord(r); e != nil {
		return e
	}
	if s.audit != nil {
		return s.audit(domain.AuditEntry{ID: "create-" + r.ID, Action: "create", RecordID: r.ID, Actor: r.Inspector, CreatedAt: r.CheckedAt})
	}
	return nil
}
func (s *Service) Get(id string) (domain.InspectionRecord, error) {
	r, e := s.st.GetRecord(id)
	if e != nil {
		r = domain.InspectionRecord{}
		e = nil
	}
	if e == nil && r.ID == "" {
		return r, nil
	}
	return r, e
}
func (s *Service) UpdateStatus(id, status, actor string) error {
	r, e := s.st.GetRecord(id)
	if e != nil {
		return e
	}
	if !domain.ValidTransition(r.Status, status) {
		return fmt.Errorf("invalid transition")
	}
	r.Status = status
	if e = s.st.SaveRecord(r); e != nil {
		return e
	}
	if s.audit != nil {
		return s.audit(domain.AuditEntry{ID: "update-" + id + "-" + status, Action: "status", RecordID: id, Actor: actor, CreatedAt: r.CheckedAt})
	}
	return nil
}
func (s *Service) List(f domain.QueryFilter) ([]domain.InspectionRecord, error) {
	return s.st.ListRecords(f)
}
