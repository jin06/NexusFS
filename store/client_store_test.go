package store

import (
	"go.etcd.io/bbolt"
	"testing"
)

func TestNewClientStore(t *testing.T) {
	_, err := NewClientStore(BoltStoreVal("test_client.db"))
	if err != nil {
		t.Error(err)
	}
}

func TestBoltStore(t *testing.T) {
	s, err := newBoltStore(Options{
		"abce.db",
	})
	if err != nil {
		t.Error(err)
	}
	err = s.Connect()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err = s.Disconnect()
		if err != nil {
			t.Error(err)
		}
	}()
	err = s.db.Update(func(tx *bbolt.Tx) error {
		bu, er := tx.CreateBucket([]byte("bucket_01"))
		if er != nil {
			t.Error(er)
		}
		er = bu.Put([]byte("hello"), []byte("world"))
		if er != nil {
			t.Error(er)
		}
		by := bu.Get([]byte("hello"))
		t.Log("bolt store get()", string(by))
		return nil
	})
	if err != nil {
		t.Error(err)
	}

}
