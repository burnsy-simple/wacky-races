package repository

import (
	"github.com/go-kit/kit/log"

	"context"
	"github.com/burnsy/wacky-races/models"
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

func NewRaceRepository(logger log.Logger) RaceRepository {
	return &mongoRepository{
		Logger: logger,
	}
}
