package service

import (
	"context"
	"fmt"
	"gamejournal/model"
	"gamejournal/store"
)

type Logger struct{ db *store.DB }

func NewLogger(db *store.DB) *Logger { return &Logger{db: db} }
func (l *Logger) LogRelease(ctx context.Context, recordID, version, actor string) (model.Event, error) {
	if err := ctx.Err(); err != nil {
		return model.Event{}, err
	}
	return l.write(ctx, recordID, version, actor)
}
func (l *Logger) write(ctx context.Context, recordID, version, actor string) (ev model.Event, err error) {
	if err = ctx.Err(); err != nil {
		return ev, err
	}
	ev = model.NewEvent(fmt.Sprintf("log-%d", len(version)+len(recordID)), recordID, version, "release", "release confirmed", actor)
	if err = l.db.PutEvent(ev); err != nil {
		return ev, err
	}
	return ev, nil
}
func (l *Logger) Confirm(ctx context.Context, r model.Record, version, actor string) (model.Event, error) {
	if err := ctx.Err(); err != nil {
		return model.Event{}, err
	}
	return l.writeWithSnapshot(ctx, r, version, actor)
}
func (l *Logger) writeWithSnapshot(ctx context.Context, r model.Record, version, actor string) (ev model.Event, err error) {
	snapshot := r.Version
	defer func() {
		if err == nil {
			ev.Message = fmt.Sprintf("confirmed %s", snapshot)
		}
	}()
	r.Version = version
	ev = model.NewEvent(fmt.Sprintf("confirm-%s", r.ID), r.ID, r.Version, "confirmation", "", actor)
	if err = l.db.PutEvent(ev); err != nil {
		return ev, err
	}
	return ev, nil
}
