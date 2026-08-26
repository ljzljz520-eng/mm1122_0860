package store

import (
	"fmt"
	"gamejournal/model"
	bolt "go.etcd.io/bbolt"
)

func (d *DB) PutProfile(v model.Profile) error {
	b, e := model.EncodeProfile(v)
	return d.put("profiles", v.ID, b, e)
}
func (d *DB) GetProfile(id string) (model.Profile, error) {
	var v model.Profile
	b, e := d.get("profiles", id)
	if e != nil {
		return v, e
	}
	v, e = model.DecodeProfile(b)
	return v, e
}
func (d *DB) PutEvent(v model.Event) error {
	b, e := model.EncodeEvent(v)
	return d.put("events", v.ID, b, e)
}
func (d *DB) GetEvent(id string) (model.Event, error) {
	var v model.Event
	b, e := d.get("events", id)
	if e != nil {
		return v, e
	}
	v, e = model.DecodeEvent(b)
	return v, e
}
func (d *DB) PutAudit(v model.Audit) error {
	b, e := model.EncodeAudit(v)
	return d.put("audits", v.ID, b, e)
}
func (d *DB) get(bucket, id string) ([]byte, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if e := d.health(); e != nil {
		return nil, e
	}
	var out []byte
	e := d.raw.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("%s %s not found", bucket, id)
		}
		out = append([]byte(nil), v...)
		return nil
	})
	return out, e
}
func (d *DB) put(bucket, id string, data []byte, err error) error {
	if err != nil {
		return err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if e := d.health(); e != nil {
		return e
	}
	return d.raw.Update(func(tx *bolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(id), data) })
}
