package models

type Meet struct {
	ID       string       `json:"id"`
	Location string       `json:"location,omitempty"`
	Category RaceCategory `json:"type"`
}

func NewMeet(id string, location string, category RaceCategory) *Meet {
	return &Meet{
		ID:       id,
		Location: location,
		Category: category,
	}
}
