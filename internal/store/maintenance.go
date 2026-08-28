package store

import (
	"database/sql"
	"inspectionbase/internal/domain"
)

func (s *Store) SaveDevices(v []domain.Device) error {
	for _, d := range v {
		if e := s.SaveDevice(d); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) LoadDevices() ([]domain.Device, error) {
	rows, e := s.db.Query(`SELECT id,name,location,active FROM devices ORDER BY id`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Device{}
	for rows.Next() {
		var d domain.Device
		var a int
		if e = rows.Scan(&d.ID, &d.Name, &d.Location, &a); e != nil {
			return nil, e
		}
		d.Active = a == 1
		out = append(out, d)
	}
	return out, rows.Err()
}
func (s *Store) SaveSettings(v []domain.InspectionSetting) error {
	for _, x := range v {
		if e := s.SaveSetting(x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) LoadSettings() ([]domain.InspectionSetting, error) {
	rows, e := s.db.Query(`SELECT key,value,updated_at FROM settings ORDER BY key`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.InspectionSetting{}
	for rows.Next() {
		var x domain.InspectionSetting
		if e = rows.Scan(&x.Key, &x.Value, &x.UpdatedAt); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (s *Store) UpdateRecordNotes(id, notes string) error {
	_, e := s.db.Exec(`UPDATE records SET notes=? WHERE id=?`, notes, id)
	return e
}
func (s *Store) UpdateInspector(id, name string) error {
	_, e := s.db.Exec(`UPDATE records SET inspector=? WHERE id=?`, name, id)
	return e
}
func (s *Store) RecordStatus(id string) (string, error) {
	var v string
	e := s.db.QueryRow(`SELECT status FROM records WHERE id=?`, id).Scan(&v)
	if e == sql.ErrNoRows {
		return "", domain.ErrNotFound
	}
	return v, e
}
func (s *Store) DeviceActive(id string) (bool, error) { d, e := s.GetDevice(id); return d.Active, e }
func (s *Store) DeleteDevice(id string) error {
	_, e := s.db.Exec(`DELETE FROM devices WHERE id=?`, id)
	return e
}
func (s *Store) DeleteAudit(id string) error {
	_, e := s.db.Exec(`DELETE FROM audits WHERE id=?`, id)
	return e
}
func (s *Store) ClearRecords() error      { _, e := s.db.Exec(`DELETE FROM records`); return e }
func (s *Store) ClearAudits() error       { _, e := s.db.Exec(`DELETE FROM audits`); return e }
func (s *Store) Backup(path string) error { _, e := s.db.Exec(`VACUUM INTO ?`, path); return e }
func (s *Store) TableNames() []string     { return []string{"devices", "records", "audits", "settings"} }
func (s *Store) SchemaReady() bool {
	for _, t := range s.TableNames() {
		if _, e := s.Count(t); e != nil {
			return false
		}
	}
	return true
}
func (s *Store) IsClosed() bool { return s.db.Ping() != nil }
func (s *Store) Stats() map[string]int {
	m := map[string]int{}
	for _, t := range s.TableNames() {
		n, e := s.Count(t)
		if e == nil {
			m[t] = n
		}
	}
	return m
}
func (s *Store) ReplaceRecord(r domain.InspectionRecord) error   { return s.SaveRecord(r) }
func (s *Store) ReplaceDevice(d domain.Device) error             { return s.SaveDevice(d) }
func (s *Store) ReplaceAudit(a domain.AuditEntry) error          { return s.SaveAudit(a) }
func (s *Store) ReplaceSetting(v domain.InspectionSetting) error { return s.SaveSetting(v) }
func (s *Store) TransactionalCreate(r domain.InspectionRecord, a domain.AuditEntry) error {
	return s.WithTransaction(func(tx *sql.Tx) error {
		if e := SaveRecordTx(tx, r); e != nil {
			return e
		}
		_, e := tx.Exec(`INSERT OR REPLACE INTO audits VALUES(?,?,?,?,?)`, a.ID, a.Action, a.RecordID, a.Actor, a.CreatedAt)
		return e
	})
}
func (s *Store) TransactionalStatus(r domain.InspectionRecord, a domain.AuditEntry) error {
	return s.TransactionalCreate(r, a)
}
