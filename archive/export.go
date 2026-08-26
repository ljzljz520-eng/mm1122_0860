package archive

import (
	"encoding/json"
	"gamejournal/model"
)

func EncodeSnapshot(r model.Record, a model.Audit) ([]byte, error) {
	return json.Marshal(struct {
		Record model.Record `json:"record"`
		Audit  model.Audit  `json:"audit"`
	}{r, a})
}
