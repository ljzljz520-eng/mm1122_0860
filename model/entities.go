package model

import "time"

type Record struct {
	ID, Title, Description, Status, Version string
	CreatedAt, UpdatedAt                    time.Time
	Tags                                    []string
}
type Profile struct {
	ID, Name, Email, Role string
	Active                bool
	CreatedAt             time.Time
}
type Event struct {
	ID, RecordID, Version, Kind, Message, Actor string
	At                                          time.Time
	Metadata                                    map[string]string
}
type Audit struct {
	ID, Entity, EntityID, Action, Actor, Detail string
	At                                          time.Time
}

func NewRecord(id, title, description string) Record {
	now := time.Now().UTC()
	return Record{ID: id, Title: title, Description: description, Status: "draft", Version: "v1", CreatedAt: now, UpdatedAt: now}
}
func (r Record) IsArchived() bool { return r.Status == "archived" }
func (r Record) CanProcess() bool { return r.Status == "draft" || r.Status == "review" }
func (r *Record) MarkProcessed(version string) {
	r.Status = "processed"
	r.Version = version
	r.UpdatedAt = time.Now().UTC()
}
func (r *Record) Archive() { r.Status = "archived"; r.UpdatedAt = time.Now().UTC() }
func NewProfile(id, name, email, role string) Profile {
	return Profile{ID: id, Name: name, Email: email, Role: role, Active: true, CreatedAt: time.Now().UTC()}
}
func NewEvent(id, recordID, version, kind, message, actor string) Event {
	return Event{ID: id, RecordID: recordID, Version: version, Kind: kind, Message: message, Actor: actor, At: time.Now().UTC(), Metadata: map[string]string{}}
}
func NewAudit(id, entity, entityID, action, actor, detail string) Audit {
	return Audit{ID: id, Entity: entity, EntityID: entityID, Action: action, Actor: actor, Detail: detail, At: time.Now().UTC()}
}
