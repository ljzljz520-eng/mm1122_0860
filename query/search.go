package query

import (
	"context"
	"gamejournal/model"
	"gamejournal/store"
	"sort"
	"strings"
)

type Searcher struct{ db *store.DB }

func New(db *store.DB) *Searcher { return &Searcher{db: db} }
func (s *Searcher) All(ctx context.Context) ([]model.Record, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	out, e := s.db.ListRecords()
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out, e
}
func (s *Searcher) ByStatus(ctx context.Context, status string) ([]model.Record, error) {
	all, e := s.All(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range all {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Searcher) Text(ctx context.Context, term string) ([]model.Record, error) {
	all, e := s.All(ctx)
	if e != nil {
		return nil, e
	}
	term = strings.ToLower(term)
	out := []model.Record{}
	for _, r := range all {
		if strings.Contains(strings.ToLower(r.Title), term) || strings.Contains(strings.ToLower(r.Description), term) {
			out = append(out, r)
		}
	}
	return out, nil
}
