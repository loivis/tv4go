package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStopstream(t *testing.T) {
	t.Run("MissingParameters", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/stopstream", nil)

		Stopstream(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("unexpected status code: %s", http.StatusText(rec.Code))
		}
	})
}
