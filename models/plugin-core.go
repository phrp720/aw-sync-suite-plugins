package models

import "time"

type Plugin interface {
	Initialize()
	Execute(events Events, watcher string, userID string, includeHostName bool) Events
	ReplicateConfig(path string)
	RawName() string
	Name() string
}

// Event represents an event in the aw database
type Event struct {
	ID        int                    `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	Duration  float64                `json:"duration"`
	Data      map[string]interface{} `json:"data"`
}

type Events []Event
