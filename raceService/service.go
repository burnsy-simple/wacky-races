package raceService

import (
	"context"
	"errors"

	"github.com/burnsy/wacky-races/repository"
	"github.com/go-kit/kit/log"

	"github.com/burnsy/wacky-races/models"
)

// Service is a simple Fetcher interface for horse/greyhound races.
type Service interface {
	GetNextRaces(ctx context.Context, num int) (models.Races, error)
	GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error)
}

var (
	ErrNotFound = errors.New("not found")
	ErrBadData  = errors.New("Malformed data")
)

type nextNService struct {
	RaceRepository repository.RaceRepository
	logger         log.Logger
}

func NewNExtNService(repository repository.RaceRepository, logger log.Logger) Service {
	return &nextNService{
		RaceRepository: repository,
		logger:         logger,
	}
}

func (s *nextNService) GetNextRaces(ctx context.Context, numRaces int) (models.Races, error) {
	return models.Races{}, nil
}

func (s *nextNService) GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error) {
	return &models.RaceDetails{}, ErrNotFound
}
