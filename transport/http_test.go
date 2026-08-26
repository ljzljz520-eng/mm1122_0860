package transport

import (
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	Health(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 204 {
		t.Fatal(w.Code)
	}
}
