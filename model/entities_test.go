package model

import "testing"

func TestRecordLifecycle(t *testing.T) {
	r := NewRecord("r1", "Puzzle", "desc")
	if !r.CanProcess() {
		t.Fatal()
	}
	r.MarkProcessed("v2")
	r.Archive()
	if !r.IsArchived() {
		t.Fatal()
	}
}
