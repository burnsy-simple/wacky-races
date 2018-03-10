package models

// Meet describes a race meeting
// swagger:model
type Meet struct {
	// required: true
	ID       string       `json:"id"`
	Location string       `json:"location,omitempty"`
	Category RaceCategory `json:"type"`
}

// NewMeet creates a new race meeting
func NewMeet(id string, location string, category RaceCategory) *Meet {
	return &Meet{
		ID:       id,
		Location: location,
		Category: category,
	}
}
