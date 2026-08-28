package audit

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/store"
)

type Service struct{ st *store.Store }

func New(st *store.Store) *Service { return &Service{st: st} }
func (s *Service) Record(a domain.AuditEntry) error {
	if a.ID == "" || a.Action == "" {
		return domain.ErrInvalid
	}
	return s.st.SaveAudit(a)
}
func (s *Service) ForRecord(id string) ([]domain.AuditEntry, error) { return s.st.ListAudits(id) }
func (s *Service) Count(id string) (int, error)                     { v, e := s.ForRecord(id); return len(v), e }
