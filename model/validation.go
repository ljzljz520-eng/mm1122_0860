package model

import "errors"

func ValidateRecord(r Record) error {
	if r.ID == "" {
		return errors.New("record id required")
	}
	if r.Title == "" {
		return errors.New("record title required")
	}
	if len(r.Title) > 200 {
		return errors.New("record title too long")
	}
	if r.Status == "" {
		return errors.New("record status required")
	}
	return nil
}
func ValidateProfile(p Profile) error {
	if p.ID == "" || p.Name == "" {
		return errors.New("profile identity required")
	}
	if p.Email == "" {
		return errors.New("profile email required")
	}
	return nil
}
func ValidateEvent(e Event) error {
	if e.ID == "" || e.RecordID == "" {
		return errors.New("event identity required")
	}
	if e.Version == "" {
		return errors.New("event version required")
	}
	return nil
}
func ValidateAudit(a Audit) error {
	if a.ID == "" || a.Entity == "" || a.EntityID == "" {
		return errors.New("audit identity required")
	}
	return nil
}
func NormalizeStatus(status string) string {
	if status == "" {
		return "draft"
	}
	switch status {
	case "draft", "review", "processed", "archived":
		return status
	default:
		return "draft"
	}
}
