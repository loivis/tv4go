package handlers

import (
	"fmt"
	"time"
)

type Event struct {
	Time    time.Time
	VideoID string
	Action  string
}

type Info struct {
	Active map[string]bool
	Events []*Event
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s at %s", e.Action, e.VideoID, e.Time.Format(time.UnixDate))
}

type response struct {
	VideoID string
	Status  string
	Message string
}
