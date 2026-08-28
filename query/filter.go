package query

import (
	"gamejournal/model"
	"time"
)

func Recent(records []model.Record, since time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if r.UpdatedAt.After(since) {
			out = append(out, r)
		}
	}
	return out
}
func Counts(records []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.Status]++
	}
	return out
}
