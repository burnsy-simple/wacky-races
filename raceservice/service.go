package raceservice

import (
	"context"
	"errors"

	"github.com/burnsy/wacky-races/models"
	"github.com/burnsy/wacky-races/repository"
	"github.com/burnsy/wacky-races/service"
	log "github.com/sirupsen/logrus"
)

const defaultNumRaces = 5

var (
	// ErrBadData is a generic error returned to this service's clients
	ErrBadData = errors.New("Malformed data")
)

type nextNService struct {
	RaceRepository repository.RaceRepository
	logger         *log.Logger
}

// NewNextNService creates a new service that will return the next N (default:5) races
func NewNextNService(repository repository.RaceRepository, logger *log.Logger) service.Service {
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
