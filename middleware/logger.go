package middleware

import (
	"context"
	"time"

	"github.com/burnsy/wacky-races/models"
	"github.com/burnsy/wacky-races/raceService"
	"github.com/go-kit/kit/log"
)

// Middleware is effectively a request level decorator
type Middleware func(raceService.Service) raceService.Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next raceService.Service) raceService.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   raceService.Service
	logger log.Logger
}

func (mw loggingMiddleware) GetNextRaces(ctx context.Context, numRaces int) (races models.Races, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetRaces", "numRaces", numRaces, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetNextRaces(ctx, numRaces)
}

func (mw loggingMiddleware) GetRaceDetails(ctx context.Context, raceID string) (rd *models.RaceDetails, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetRaceDetails", "raceID", raceID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetRaceDetails(ctx, raceID)
}
