// Package models represents the domain as the user is likely to want
// to interact with it given current limited understanding.
package models

import (
	"time"
)

// Race represents a single race of any type.
// ID should be globally unique.
type Race struct {
	ID       string       `json:"id"`
	MeetID   string       `json:"meet_id"`
	Category RaceCategory `json:"type"`
	Name     string       `json:"name"`
	StartAt  time.Time    `json:"start_at"`
	CloseAt  time.Time    `json:"close_at"`
}

type Races []*Race

// Implementation of sort.Interface
func (races Races) Len() int {
	return len(races)
}
func (races Races) Swap(i, j int) {
	races[i], races[j] = races[j], races[i]
}
func (races Races) Less(i, j int) bool {
	return races[i].CloseAt.Before(races[j].CloseAt)
}

// RaceDetails represents a single race of any type.
type RaceDetails struct {
	*Race
	// Additional race specific details can go here
	Competitors []Competitor
}

func newRace(id, meetID string, category RaceCategory, name string, start, close time.Time) *Race {
	return &Race{
		ID:       id,
		MeetID:   meetID,
		Category: category,
		Name:     name,
		StartAt:  start,
		CloseAt:  close,
	}
}

func NewThoroughbredRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Thoroughbred, name, start, close)
}

func NewGreyhoundRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Greyhound, name, start, close)
}

func NewHarnessRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Harness, name, start, close)
}
