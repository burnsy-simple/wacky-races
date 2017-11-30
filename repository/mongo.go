// Package repository is our abstraction of the data source.
// In the context of this problem we might model one aggregate for a race with the competitors contained within
// and have another aggregate for the Meets. These aggregate could be stored in 2 separate collections in e.g. Mongo
// (or we could have multiple table in SQL).
// In the real world this data might be fed in to us from external services...
package repository

import (
	"github.com/burnsy/wacky-races/models"
	"golang.org/x/net/context"
)

func (mr mongoRepository) GetNextNRaces(ctx context.Context, numRaces int) (models.Races, error) {
	return models.Races{}, nil
}

func (mr mongoRepository) GetRaceDetails(ctx context.Context, id string) (*models.RaceDetails, error) {
	return &models.RaceDetails{}, nil
}
