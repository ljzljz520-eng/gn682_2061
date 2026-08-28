package store

import (
	"database/sql"
	"inspectionbase/internal/domain"
)

func (s *Store) WithTransaction(fn func(*sql.Tx) error) error {
	tx, e := s.db.Begin()
	if e != nil {
		return e
	}
	if e = fn(tx); e != nil {
		tx.Rollback()
		return e
	}
	return tx.Commit()
}
func SaveRecordTx(tx *sql.Tx, r domain.InspectionRecord) error {
	_, e := tx.Exec(`INSERT OR REPLACE INTO records VALUES(?,?,?,?,?,?)`, r.ID, r.DeviceID, r.Status, r.Inspector, r.Notes, r.CheckedAt)
	return e
}
func (s *Store) DeleteRecord(id string) error {
	_, e := s.db.Exec(`DELETE FROM records WHERE id=?`, id)
	return e
}
func (s *Store) Exists(id string) bool {
	var n int
	e := s.db.QueryRow(`SELECT count(1) FROM records WHERE id=?`, id).Scan(&n)
	return e == nil && n > 0
}
