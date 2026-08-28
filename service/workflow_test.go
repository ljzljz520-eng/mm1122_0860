package service

import (
	"context"
	"gamejournal/model"
	"gamejournal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "a"))
	defer d.Close()
	j := New(d)
	if e := j.Register(context.Background(), model.NewRecord("r", "Title", "desc")); e != nil {
		t.Fatal(e)
	}
	if e := j.Review(context.Background(), "r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "b"))
	defer d.Close()
	j := New(d)
	j.Register(context.Background(), model.NewRecord("r", "Title", "desc"))
	j.Review(context.Background(), "r")
	if _, e := j.Process(context.Background(), "r", "v2", "dev"); e != nil {
		t.Fatal(e)
	}
	if e := j.Archive(context.Background(), "r", "dev"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "c"))
	defer d.Close()
	j := New(d)
	r := model.NewRecord("r", "Title", "desc")
	j.Register(context.Background(), r)
	l := NewLogger(d)
	if _, e := l.Confirm(context.Background(), r, "v2", "dev"); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain15(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer d.Close()
	j := New(d)
	if e := j.Register(context.Background(), model.NewRecord("chain", "Chain", "track")); e != nil {
		t.Fatal(e)
	}
}
