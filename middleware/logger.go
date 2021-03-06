package middleware

import (
	"context"
	"time"

	"github.com/burnsy/wacky-races/models"
	"github.com/burnsy/wacky-races/service"
	log "github.com/sirupsen/logrus"
)

// Middleware is effectively a request level decorator
type Middleware func(service.Service) service.Service

// LoggingMiddleware creates the logging middleware
func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   service.Service
	logger *log.Logger
}

func (mw loggingMiddleware) GetNextRaces(ctx context.Context, numRaces int) (races models.Races, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{
			"numRaces": numRaces,
			"took":     time.Since(begin),
			"err":      err,
		}).Infof("GetRaces")
	}(time.Now())
	return mw.next.GetNextRaces(ctx, numRaces)
}

func (mw loggingMiddleware) GetRaceDetails(ctx context.Context, raceID string) (rd *models.RaceDetails, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(log.Fields{
			"raceID": raceID,
			"took":   time.Since(begin),
			"err":    err,
		}).Infof("GetRaceDetails")
	}(time.Now())
	return mw.next.GetRaceDetails(ctx, raceID)
}
