package model

import "testing"

func TestValidation(t *testing.T) {
	if ValidateRecord(Record{}) == nil {
		t.Fatal()
	}
	if ValidateProfile(Profile{}) == nil {
		t.Fatal()
	}
	if ValidateEvent(Event{}) == nil {
		t.Fatal()
	}
	if ValidateAudit(Audit{}) == nil {
		t.Fatal()
	}
}
