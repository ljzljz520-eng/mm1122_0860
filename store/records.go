package store

import (
	"fmt"
	"gamejournal/model"
	bolt "go.etcd.io/bbolt"
)

func (d *DB) PutRecord(r model.Record) error {
	b, err := model.EncodeRecord(r)
	if err != nil {
		return err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err = d.health(); err != nil {
		return err
	}
	return d.raw.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte("records")).Put([]byte(r.ID), b) })
}
func (d *DB) GetRecord(id string) (model.Record, error) {
	var r model.Record
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.health(); err != nil {
		return r, err
	}
	err := d.raw.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("records")).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("record %s not found", id)
		}
		var e error
		r, e = model.DecodeRecord(v)
		return e
	})
	return r, err
}
func (d *DB) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.health(); err != nil {
		return nil, err
	}
	err := d.raw.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			r, e := model.DecodeRecord(v)
			if e == nil {
				out = append(out, r)
			}
			return e
		})
	})
	return out, err
}
func (d *DB) DeleteRecord(id string) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.health(); err != nil {
		return err
	}
	return d.raw.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
