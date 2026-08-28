package audit

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/store"
)

type Query struct {
	Action, Actor string
	Limit         int
}

func NewQuery() Query { return Query{Limit: 100} }
func (q Query) Match(a domain.AuditEntry) bool {
	if q.Action != "" && q.Action != a.Action {
		return false
	}
	if q.Actor != "" && q.Actor != a.Actor {
		return false
	}
	return true
}
func Filter(v []domain.AuditEntry, q Query) []domain.AuditEntry {
	out := []domain.AuditEntry{}
	for _, a := range v {
		if q.Match(a) {
			out = append(out, a)
			if q.Limit > 0 && len(out) >= q.Limit {
				break
			}
		}
	}
	return out
}
func ActionCounts(v []domain.AuditEntry) map[string]int {
	m := map[string]int{}
	for _, a := range v {
		m[a.Action]++
	}
	return m
}
func Actors(v []domain.AuditEntry) map[string]int {
	m := map[string]int{}
	for _, a := range v {
		m[a.Actor]++
	}
	return m
}
func Latest(v []domain.AuditEntry) domain.AuditEntry {
	if len(v) == 0 {
		return domain.AuditEntry{}
	}
	return v[len(v)-1]
}
func Contains(v []domain.AuditEntry, id string) bool {
	for _, a := range v {
		if a.ID == id {
			return true
		}
	}
	return false
}
func ValidAll(v []domain.AuditEntry) bool {
	for _, a := range v {
		if !domain.AuditActionValid(a) {
			return false
		}
	}
	return true
}
func LoadForRecord(s *store.Store, id string) ([]domain.AuditEntry, error) { return s.ListAudits(id) }
func CountAction(v []domain.AuditEntry, action string) int {
	n := 0
	for _, a := range v {
		if a.Action == action {
			n++
		}
	}
	return n
}
func HasAction(v []domain.AuditEntry, action string) bool { return CountAction(v, action) > 0 }
