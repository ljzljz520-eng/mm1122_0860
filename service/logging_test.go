package service

import (
	"context"
	"gamejournal/model"
	"gamejournal/store"
	"path/filepath"
	"testing"
)

func TestConfirmVersions(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "l"))
	defer d.Close()
	l := NewLogger(d)
	r := model.NewRecord("r", "T", "D")
	a, _ := l.Confirm(context.Background(), r, "v2", "a")
	b, _ := l.Confirm(context.Background(), r, "v3", "a")
	if a.Version != "v2" || b.Version != "v3" {
		t.Fatal()
	}
	if b.Message != "confirmed v3" {
		t.Fatalf("message=%q", b.Message)
	}
}
