package store

func (s *Store) Ping() error { return s.db.Ping() }
func (s *Store) Count(table string) (int, error) {
	var n int
	e := s.db.QueryRow("SELECT count(1) FROM " + table).Scan(&n)
	return n, e
}
func (s *Store) Vacuum() error { _, e := s.db.Exec("VACUUM"); return e }
