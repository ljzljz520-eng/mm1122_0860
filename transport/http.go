package transport

import (
	"context"
	"encoding/json"
	"gamejournal/model"
	"gamejournal/service"
	"net/http"
)

type Server struct{ journal *service.Journal }

func New(j *service.Journal) *Server { return &Server{journal: j} }
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/records", s.records)
	return mux
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method", 405)
		return
	}
	var rec model.Record
	if json.NewDecoder(r.Body).Decode(&rec) != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if err := s.journal.Register(context.Background(), rec); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
