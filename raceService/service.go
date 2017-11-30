package raceService

import (
	"context"
	"errors"

	"github.com/burnsy/wacky-races/repository"
	"github.com/go-kit/kit/log"

	"github.com/burnsy/wacky-races/models"
)

const defaultNumRaces = 5

// Service is a simple Fetcher interface for horse/greyhound races.
type Service interface {
	GetNextRaces(ctx context.Context, num int) (models.Races, error)
	GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error)
}

var (
	ErrBadData = errors.New("Malformed data")
)

type nextNService struct {
	RaceRepository repository.RaceRepository
	logger         log.Logger
}

func NewNextNService(repository repository.RaceRepository, logger log.Logger) Service {
	return &nextNService{
		RaceRepository: repository,
		logger:         logger,
	}
}

func (svc *nextNService) GetNextRaces(ctx context.Context, numRaces int) (models.Races, error) {
	if numRaces < 0 {
		return nil, ErrBadData
	} else if numRaces == 0 {
		numRaces = defaultNumRaces
	}
	return svc.RaceRepository.GetNextNRaces(ctx, numRaces)
}

func (svc *nextNService) GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error) {
	if len(raceID) == 0 {
		return nil, ErrBadData
	}

	return svc.RaceRepository.GetRaceDetails(ctx, raceID)
}
