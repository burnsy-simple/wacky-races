package repository

import (
	"context"

	"github.com/burnsy/wacky-races/models"
	"github.com/go-kit/kit/log"
)

// RaceRepository is our database abstraction
// For now, we don't connect to a real database
type RaceRepository interface {
	// GetNextNRaces returns the next N races
	GetNextNRaces(ctx context.Context, numRaces int) (models.Races, error)
	// GetRaceDetails returns the race details
	GetRaceDetails(ctx context.Context, id string) (*models.RaceDetails, error)
}

// We won't really call a Mongo DB for now but we could...
type mongoRepository struct {
	RaceRepository
	Logger log.Logger
}

// NewRaceRepository creates a new repository
func NewRaceRepository(logger log.Logger) RaceRepository {
	return &mongoRepository{
		Logger: logger,
	}
}
