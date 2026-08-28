package service

import (
	"context"
	"fmt"
	"gamejournal/model"
	"gamejournal/store"
	"sync"
	"time"
)

type Journal struct {
	db       *store.DB
	mu       sync.Mutex
	sequence int
}

func New(db *store.DB) *Journal { return &Journal{db: db} }
func (j *Journal) Register(ctx context.Context, r model.Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := model.ValidateRecord(r); err != nil {
		return err
	}
	r.Status = model.NormalizeStatus(r.Status)
	return j.db.PutRecord(r)
}
func (j *Journal) CreateProfile(ctx context.Context, p model.Profile) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := model.ValidateProfile(p); err != nil {
		return err
	}
	return j.db.PutProfile(p)
}
func (j *Journal) Review(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r, e := j.db.GetRecord(id)
	if e != nil {
		return e
	}
	if r.IsArchived() {
		return fmt.Errorf("archived record")
	}
	r.Status = "review"
	r.UpdatedAt = time.Now().UTC()
	return j.db.PutRecord(r)
}
func (j *Journal) Process(ctx context.Context, id, version, actor string) (model.Event, error) {
	if err := ctx.Err(); err != nil {
		return model.Event{}, err
	}
	j.mu.Lock()
	j.sequence++
	seq := j.sequence
	j.mu.Unlock()
	r, e := j.db.GetRecord(id)
	if e != nil {
		return model.Event{}, e
	}
	if !r.CanProcess() {
		return model.Event{}, fmt.Errorf("record not processable")
	}
	r.MarkProcessed(version)
	if e = j.db.PutRecord(r); e != nil {
		return model.Event{}, e
	}
	ev := model.NewEvent(fmt.Sprintf("event-%d", seq), id, version, "release", "processed", actor)
	if e = j.db.PutEvent(ev); e != nil {
		return ev, e
	}
	return ev, nil
}
func (j *Journal) Archive(ctx context.Context, id, actor string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r, e := j.db.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "processed" {
		return fmt.Errorf("record must be processed")
	}
	r.Archive()
	if e = j.db.PutRecord(r); e != nil {
		return e
	}
	return j.db.PutAudit(model.NewAudit(fmt.Sprintf("audit-%d", time.Now().UnixNano()), "Record", id, "archive", actor, "archived"))
}
func (j *Journal) Fetch(ctx context.Context, id string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return model.Record{}, err
	}
	return j.db.GetRecord(id)
}
