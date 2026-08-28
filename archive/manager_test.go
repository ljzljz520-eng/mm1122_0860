package archive

import (
	"context"
	"gamejournal/model"
	"gamejournal/store"
	"path/filepath"
	"testing"
)

func TestArchiveRestore(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "z"))
	defer d.Close()
	r := model.NewRecord("r", "T", "D")
	r.MarkProcessed("v2")
	r.Archive()
	d.PutRecord(r)
	m := New(d)
	if _, e := m.Snapshot(context.Background(), "r"); e != nil {
		t.Fatal(e)
	}
	if _, e := m.Restore(context.Background(), "r"); e != nil {
		t.Fatal(e)
	}
}
