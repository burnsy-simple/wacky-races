// Package models represents the domain as the user is likely to want
// to interact with it given current limited understanding.
package models

import (
	"time"
)

// Race represents a single race of any type.
// ID should be globally unique.
// swagger:model race
type Race struct {
	// required: true
	ID string `json:"id"`
	// required: true
	MeetID string `json:"meet_id"`
	// required: true
	Category RaceCategory `json:"type"`
	// required: true
	Name string `json:"name"`
	// required: true
	StartAt time.Time `json:"start_at"`
	// required: true
	CloseAt time.Time `json:"close_at"`
}

// Races dscribes a bunch of races
// swagger:model races
type Races []Race

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
// swagger:model raceDetails
type RaceDetails struct {
	// swagger:allOf
	// required: true
	Race *Race

	// required: true
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

// NewThoroughbredRace creates a thoroughbred/horse race
func NewThoroughbredRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Thoroughbred, name, start, close)
}

// NewGreyhoundRace creates a greyhound race
func NewGreyhoundRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Greyhound, name, start, close)
}

// NewHarnessRace creates a harness race
func NewHarnessRace(id, meetID string, name string, start, close time.Time) *Race {
	return newRace(id, meetID, Harness, name, start, close)
}
