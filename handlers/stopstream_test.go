package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserInfoStore struct {
}

func (s *mockUserInfoStore) Stop(userID, videoID string) error {
	if userID != "123" || videoID != "456" {
		return fmt.Errorf("unexpected input: userID %q, videoID %q", userID, videoID)
	}
	return nil
}

func TestStopstream(t *testing.T) {
	t.Run("StoppingStartedStream", func(t *testing.T) {
		var store mockUserInfoStore

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/stopstream?site=cmore.se&user_id=123&video_id=456", nil)

		srv := &Server{Store: &store}
		srv.Stopstream(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("unexpected status code: %s", http.StatusText(rec.Code))
		}
	})

	t.Run("MissingParameters", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/stopstream", nil)

		srv := &Server{}
		srv.Stopstream(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("unexpected status code: %s", http.StatusText(rec.Code))
		}
	})
}
