package server

import "go.etcd.io/bbolt"

type Manager struct {
	DB *bbolt.DB
}

func New(db_path string) (m *Manager, err error) {
	m = &Manager{}
	m.DB, err = bbolt.Open(db_path, 0600, nil)
	if err != nil {
		return
	}
	return
}
