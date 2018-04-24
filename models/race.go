// Package models represents the domain as the user is likely to want
// to interact with it given current limited understanding.
package models

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// Race represents a single race of any type.
// ID should be globally unique.
// swagger:model race
type Race struct {
	// required: true
	ID string `json:"id" validate:"required"`
	// required: true
	MeetID string `json:"meet_id" validate:"required"`
	// required: true
	Category RaceCategory `json:"type" validate:"gte=0,lte=2"`
	// required: true
	Name string `json:"name" validate:"required"`
	// required: true
	StartAt time.Time `json:"start_at" validate:"required"`
	// required: true
	CloseAt time.Time `json:"close_at" validate:"required"`
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
	// required: true
	Race *Race

	// required: true
	Competitors []Competitor
}

func newRace(id, meetID string, category RaceCategory, name string, start, close time.Time) *Race {
	validate := validator.New()

	race := &Race{
		ID:       id,
		MeetID:   meetID,
		Category: category,
		Name:     name,
		StartAt:  start,
		CloseAt:  close,
	}
	err := validate.Struct(race)
	if err != nil {
		log.Infof("got err: %v", err)
		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}

	return race
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
