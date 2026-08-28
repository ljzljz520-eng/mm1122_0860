package service

import "gamejournal/model"

func AllowedTransition(from, to string) bool {
	switch from {
	case "draft":
		return to == "review"
	case "review":
		return to == "processed"
	case "processed":
		return to == "archived"
	case "archived":
		return to == "processed"
	}
	return false
}
func Summarize(r model.Record) string { return r.ID + ":" + r.Status + ":" + r.Version }
