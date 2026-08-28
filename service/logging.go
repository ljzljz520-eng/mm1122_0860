package service

import (
	"context"
	"fmt"
	"gamejournal/model"
	"gamejournal/store"
	"sync"
)

type Logger struct {
	db       *store.DB
	mu       sync.Mutex
	sequence int
}

func NewLogger(db *store.DB) *Logger { return &Logger{db: db} }

// nextID issues a monotonically unique event id so that consecutive log
// entries for the same record do not collide and overwrite one another.
func (l *Logger) nextID(prefix string) string {
	l.mu.Lock()
	l.sequence++
	seq := l.sequence
	l.mu.Unlock()
	return fmt.Sprintf("%s-%d", prefix, seq)
}

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
	ev = model.NewEvent(l.nextID("log"), recordID, version, "release", "release confirmed", actor)
	if err = l.db.PutEvent(ev); err != nil {
		return ev, err
	}
	return ev, nil
}
func (l *Logger) Confirm(ctx context.Context, r model.Record, version, actor string) (model.Event, error) {
	if err := ctx.Err(); err != nil {
		return model.Event{}, err
	}
	return l.writeConfirm(ctx, r, version, actor)
}
func (l *Logger) writeConfirm(ctx context.Context, r model.Record, version, actor string) (ev model.Event, err error) {
	if err = ctx.Err(); err != nil {
		return ev, err
	}
	// Each log entry records its own version. Build the full event, including
	// the message, before persisting so the stored entry reflects this release
	// rather than a stale snapshot carried over from a previous one.
	r.Version = version
	ev = model.NewEvent(l.nextID("confirm"), r.ID, r.Version, "confirmation", fmt.Sprintf("confirmed %s", r.Version), actor)
	if err = l.db.PutEvent(ev); err != nil {
		return ev, err
	}
	return ev, nil
}
