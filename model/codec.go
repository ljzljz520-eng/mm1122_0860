package model

import "encoding/json"

func EncodeRecord(v Record) ([]byte, error) { return json.Marshal(v) }
func DecodeRecord(b []byte) (Record, error) {
	var v Record
	err := json.Unmarshal(b, &v)
	return v, err
}
func EncodeProfile(v Profile) ([]byte, error) { return json.Marshal(v) }
func DecodeProfile(b []byte) (Profile, error) {
	var v Profile
	err := json.Unmarshal(b, &v)
	return v, err
}
func EncodeEvent(v Event) ([]byte, error) { return json.Marshal(v) }
func DecodeEvent(b []byte) (Event, error) { var v Event; err := json.Unmarshal(b, &v); return v, err }
func EncodeAudit(v Audit) ([]byte, error) { return json.Marshal(v) }
func DecodeAudit(b []byte) (Audit, error) { var v Audit; err := json.Unmarshal(b, &v); return v, err }
