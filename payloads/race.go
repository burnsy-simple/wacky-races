package payloads

import "github.com/burnsy/wacky-races/models"

// It's often better to return service (business logic) errors in the
// response object. This means we have to do a bit more work in the HTTP
// response encoder to detect e.g. a not-found error and provide a proper HTTP
// status code. That work is done with the errorer interface, in transport.go.
// Response types that may contain business-logic errors implement that
// interface.

type RacesReq struct {
	NumRaces int
}

// RacesResp represents a single race of any type.
// ID should be globally unique.
type RacesResp struct {
	models.Races
	Err error
}

func (r RacesResp) error() error { return r.Err }

type RaceDetailsReq struct {
	RaceID string
}

// RaceDetailsResp represents a single race of any type.
type RaceDetailsResp struct {
	*models.RaceDetails
	Err error
}

func (r RaceDetailsResp) error() error { return r.Err }
