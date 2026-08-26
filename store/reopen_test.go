package store

import (
	"gamejournal/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "journal.db")
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("persist", "Persist", "x")
	if e = d.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	d.Close()
	d, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	got, e := d.GetRecord("persist")
	if e != nil || got.Title != "Persist" {
		t.Fatalf("%v %+v", e, got)
	}
}
