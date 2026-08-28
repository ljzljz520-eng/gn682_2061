package store

import (
	"database/sql"
	"inspectionbase/internal/domain"
	_ "modernc.org/sqlite"
)

type Store struct {
	db   *sql.DB
	path string
}

func Open(path string) (*Store, error) {
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	if e = s.init(); e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) init() error {
	_, e := s.db.Exec(`CREATE TABLE IF NOT EXISTS devices(id TEXT PRIMARY KEY,name TEXT,location TEXT,active INTEGER); CREATE TABLE IF NOT EXISTS records(id TEXT PRIMARY KEY,device_id TEXT,status TEXT,inspector TEXT,notes TEXT,checked_at TEXT); CREATE TABLE IF NOT EXISTS audits(id TEXT PRIMARY KEY,action TEXT,record_id TEXT,actor TEXT,created_at TEXT); CREATE TABLE IF NOT EXISTS settings(key TEXT PRIMARY KEY,value TEXT,updated_at TEXT)`)
	return e
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Path() string { return s.path }
func (s *Store) SaveDevice(d domain.Device) error {
	_, e := s.db.Exec(`INSERT OR REPLACE INTO devices VALUES(?,?,?,?)`, d.ID, d.Name, d.Location, d.Active)
	return e
}
func (s *Store) GetDevice(id string) (domain.Device, error) {
	var d domain.Device
	var a int
	e := s.db.QueryRow(`SELECT id,name,location,active FROM devices WHERE id=?`, id).Scan(&d.ID, &d.Name, &d.Location, &a)
	d.Active = a == 1
	if e == sql.ErrNoRows {
		return d, domain.ErrNotFound
	}
	return d, e
}
