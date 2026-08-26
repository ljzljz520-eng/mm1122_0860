package archive

import (
	"context"
	"fmt"
	"gamejournal/model"
	"gamejournal/store"
	"time"
)

type Manager struct{ db *store.DB }

func New(db *store.DB) *Manager { return &Manager{db: db} }
func (m *Manager) Snapshot(ctx context.Context, id string) (model.Audit, error) {
	if err := ctx.Err(); err != nil {
		return model.Audit{}, err
	}
	r, e := m.db.GetRecord(id)
	if e != nil {
		return model.Audit{}, e
	}
	if !r.IsArchived() {
		return model.Audit{}, fmt.Errorf("record not archived")
	}
	a := model.NewAudit(fmt.Sprintf("snapshot-%d", time.Now().UnixNano()), "Record", id, "snapshot", "system", r.Version)
	return a, m.db.PutAudit(a)
}
func (m *Manager) Restore(ctx context.Context, id string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return model.Record{}, err
	}
	r, e := m.db.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !r.IsArchived() {
		return r, fmt.Errorf("record not archived")
	}
	r.Status = "processed"
	r.UpdatedAt = time.Now().UTC()
	return r, m.db.PutRecord(r)
}
func (m *Manager) Purge(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return m.db.DeleteRecord(id)
}
