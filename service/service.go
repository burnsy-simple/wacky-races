package service

import (
	"context"

	"github.com/burnsy/wacky-races/models"
)

// Service is a simple Fetcher interface for horse/greyhound races.
type Service interface {
	GetNextRaces(ctx context.Context, num int) (models.Races, error)
	GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error)
}
