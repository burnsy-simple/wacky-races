package payloads

import "github.com/burnsy/wacky-races/models"

// It's often better to return service (business logic) errors in the
// response object. This means we have to do a bit more work in the HTTP
// response encoder to detect e.g. a not-found error and provide a proper HTTP
// status code. That work is done with the errorer interface, in transport.go.
// Response types that may contain business-logic errors implement that
// interface.

// RacesReq ...
// swagger:parameters listRaces
type RacesReq struct {
	// in: path
	NumRaces int
}

// RacesResp represents a single race of any type.
// ID should be globally unique.
// swagger:response racesResponse
type RacesResp struct {
	// in: body
	Races models.Races `json:"races"`

	// in: body
	Err error `json:"err"`
}

func (r RacesResp) error() error { return r.Err }

// RaceDetailsReq ...
// swagger:parameters getRaceDetails
type RaceDetailsReq struct {
	// in: path
	RaceID string
}

// RaceDetailsResp represents a single race of any type.
// swagger:response raceDetailsResponse
type RaceDetailsResp struct {
	// in: body
	// swagger:allOf
	// required: true
	*models.RaceDetails `json:"race_details"`

	// in: body
	// required: true
	Err error `json:"err"`
}

func (r RaceDetailsResp) error() error { return r.Err }
