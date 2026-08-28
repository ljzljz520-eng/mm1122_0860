package query

import (
	"context"
	"gamejournal/model"
	"gamejournal/store"
	"path/filepath"
	"testing"
)

func TestSearch(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "q"))
	defer d.Close()
	d.PutRecord(model.NewRecord("r", "Moon Puzzle", "find me"))
	s := New(d)
	got, e := s.Text(context.Background(), "moon")
	if e != nil || len(got) != 1 {
		t.Fatal(e, len(got))
	}
}
