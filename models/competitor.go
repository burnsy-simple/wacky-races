package models

// There are three types of distinct races/competitors
// We assume competitors can only partake in one type of race
const (
	Thoroughbred = iota
	Greyhound
	Harness
)

// RaceCategory is the type of race - see above enum/const
type RaceCategory int8

// Competitor competes in a race
// ID should be unique within the race (at a minimum).
// Assumption: A competitor only competes in one type/category of race
// swagger:model competitor
type Competitor struct {
	// required: true
	ID string `json:"id"`
	// required: true
	Name string `json:"name,omitempty"`
	// required: true
	Position int8 `json:"position"`
	// required: true
	Category RaceCategory `json:"type"` // type is the domain term?
}

// NewCompetitor creates a new competitor. Ordinarily a DB would be responsible
// for creating the ID but we're hard-coding the data for now...
func NewCompetitor(id, name string, category RaceCategory) *Competitor {
	return &Competitor{
		ID:       id,
		Name:     name,
		Category: category,
	}
}

// NewThoroughbred creates a thoroughbred
func NewThoroughbred(id, name string) *Competitor {
	return NewCompetitor(id, name, Thoroughbred)
}

// NewGreyhound creates a thoroughbred
func NewGreyhound(id, name string) *Competitor {
	return NewCompetitor(id, name, Greyhound)
}

// NewHarness creates a thoroughbred
func NewHarness(id, name string) *Competitor {
	return NewCompetitor(id, name, Harness)
}
