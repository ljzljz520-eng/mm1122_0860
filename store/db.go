package store

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"time"
)

var buckets = []string{"records", "profiles", "events", "audits"}

type DB struct {
	mu   sync.RWMutex
	raw  *bolt.DB
	path string
}

func Open(path string) (*DB, error) {
	if path == "" {
		return nil, fmt.Errorf("path required")
	}
	db, err := bolt.Open(filepath.Clean(path), 0600, &bolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}
	d := &DB{raw: db, path: path}
	err = db.Update(func(tx *bolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return d, nil
}
func (d *DB) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.raw == nil {
		return nil
	}
	err := d.raw.Close()
	d.raw = nil
	return err
}
func (d *DB) Path() string { return d.path }
func (d *DB) health() error {
	if d.raw == nil {
		return fmt.Errorf("database closed")
	}
	return nil
}
