package workflow

import (
	"context"
	"gamejournal/archive"
	"gamejournal/model"
	"gamejournal/query"
	"gamejournal/service"
)

type Engine struct {
	journal *service.Journal
	search  *query.Searcher
	archive *archive.Manager
}

func New(j *service.Journal, s *query.Searcher, a *archive.Manager) *Engine {
	return &Engine{journal: j, search: s, archive: a}
}
func (e *Engine) Intake(ctx context.Context, r model.Record) error {
	if err := e.journal.Register(ctx, r); err != nil {
		return err
	}
	return e.journal.Review(ctx, r.ID)
}
func (e *Engine) Release(ctx context.Context, id, version, actor string) (model.Event, error) {
	ev, err := e.journal.Process(ctx, id, version, actor)
	if err != nil {
		return ev, err
	}
	if err = e.journal.Archive(ctx, id, actor); err != nil {
		return ev, err
	}
	return ev, nil
}
func (e *Engine) Track(ctx context.Context, term string) ([]model.Record, error) {
	return e.search.Text(ctx, term)
}
func (e *Engine) Reopen(ctx context.Context, id string) (model.Record, error) {
	return e.archive.Restore(ctx, id)
}
