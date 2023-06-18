package store

import "testing"

func TestOptions(t *testing.T) {
	value := "test_options.db"
	option := BoltStoreVal(value)
	opts := NewOptions(option)
	if opts.BoltStorePath != value {
		t.Fatalf("%s != %s", opts.BoltStorePath, value)
	}
}
