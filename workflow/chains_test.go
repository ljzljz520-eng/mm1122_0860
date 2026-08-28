package workflow

import (
	"context"
	"gamejournal/archive"
	"gamejournal/model"
	"gamejournal/query"
	"gamejournal/service"
	"gamejournal/store"
	"path/filepath"
	"testing"
)

func TestEngineChains(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "w"))
	defer d.Close()
	j := service.New(d)
	e := New(j, query.New(d), archive.New(d))
	r := model.NewRecord("r", "Track", "desc")
	if err := e.Intake(context.Background(), r); err != nil {
		t.Fatal(err)
	}
	if _, err := e.Release(context.Background(), "r", "v2", "dev"); err != nil {
		t.Fatal(err)
	}
}
