package store

import (
	"inspectionbase/internal/domain"
)

func (s *Store) SaveAudit(a domain.AuditEntry) error {
	_, e := s.db.Exec(`INSERT OR REPLACE INTO audits VALUES(?,?,?,?,?)`, a.ID, a.Action, a.RecordID, a.Actor, a.CreatedAt)
	return e
}
func (s *Store) ListAudits(recordID string) ([]domain.AuditEntry, error) {
	rows, e := s.db.Query(`SELECT id,action,record_id,actor,created_at FROM audits WHERE record_id=? ORDER BY id`, recordID)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.AuditEntry{}
	for rows.Next() {
		var a domain.AuditEntry
		if e = rows.Scan(&a.ID, &a.Action, &a.RecordID, &a.Actor, &a.CreatedAt); e != nil {
			return nil, e
		}
		out = append(out, a)
	}
	return out, rows.Err()
}
func (s *Store) SaveSetting(v domain.InspectionSetting) error {
	_, e := s.db.Exec(`INSERT OR REPLACE INTO settings VALUES(?,?,?)`, v.Key, v.Value, v.UpdatedAt)
	return e
}
func (s *Store) GetSetting(k string) (domain.InspectionSetting, error) {
	var v domain.InspectionSetting
	e := s.db.QueryRow(`SELECT key,value,updated_at FROM settings WHERE key=?`, k).Scan(&v.Key, &v.Value, &v.UpdatedAt)
	return v, e
}
