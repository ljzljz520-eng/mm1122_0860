package workflow

import (
	"context"
	"gamejournal/model"
	"gamejournal/query"
)

func BuildReport(ctx context.Context, s *query.Searcher) (map[string]interface{}, error) {
	all, e := s.All(ctx)
	if e != nil {
		return nil, e
	}
	return map[string]interface{}{"total": len(all), "records": all}, nil
}
func Versions(records []model.Record) []string {
	out := []string{}
	for _, r := range records {
		out = append(out, r.Version)
	}
	return out
}
