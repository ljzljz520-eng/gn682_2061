package store

import (
	"database/sql"
	"inspectionbase/internal/domain"
)

func (s *Store) SaveRecord(r domain.InspectionRecord) error {
	_, e := s.db.Exec(`INSERT OR REPLACE INTO records VALUES(?,?,?,?,?,?)`, r.ID, r.DeviceID, r.Status, r.Inspector, r.Notes, r.CheckedAt)
	return e
}
func (s *Store) GetRecord(id string) (domain.InspectionRecord, error) {
	var r domain.InspectionRecord
	e := s.db.QueryRow(`SELECT id,device_id,status,inspector,notes,checked_at FROM records WHERE id=?`, id).Scan(&r.ID, &r.DeviceID, &r.Status, &r.Inspector, &r.Notes, &r.CheckedAt)
	if e == sql.ErrNoRows {
		return r, domain.ErrNotFound
	}
	return r, e
}
func (s *Store) ListRecords(f domain.QueryFilter) ([]domain.InspectionRecord, error) {
	q := `SELECT id,device_id,status,inspector,notes,checked_at FROM records ORDER BY id LIMIT ?`
	rows, e := s.db.Query(q, domain.NormalizeLimit(f.Limit))
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.InspectionRecord{}
	for rows.Next() {
		var r domain.InspectionRecord
		if e = rows.Scan(&r.ID, &r.DeviceID, &r.Status, &r.Inspector, &r.Notes, &r.CheckedAt); e != nil {
			return nil, e
		}
		if f.Matches(r) {
			out = append(out, r)
		}
	}
	return out, rows.Err()
}
