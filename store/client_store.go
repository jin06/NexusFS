package store

import (
	"go.etcd.io/bbolt"
)

type ClientStore interface {
	Connect() error
	Disconnect() error
}

func NewClientStore(options ...Option) (ClientStore, error) {
	opts := NewOptions(options...)

	s, err := newBoltStore(opts)
	if err != nil {
		return nil, err
	}
	err = s.Connect()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func newBoltStore(opts Options) (*boltStore, error) {
	s := &boltStore{
		Options: opts,
	}
	return s, nil
}

type boltStore struct {
	Options
	db *bbolt.DB
}

func (s *boltStore) Connect() error {
	db, err := bbolt.Open(s.Options.BoltStorePath, 0600, nil)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *boltStore) Disconnect() error {
	return s.db.Close()
}
