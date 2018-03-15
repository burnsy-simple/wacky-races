// Package repository is our abstraction of the data source.
// In the context of this problem we might model one aggregate for a race with the competitors contained within
// and have another aggregate for the Meets. These aggregate could be stored in 2 separate collections e.g. in Mongo
// (or we could have multiple table in SQL).
// We might have the original published race details in our DB with a feed from external services providing updates..
package repository

import (
	"sort"
	"time"

	"context"
	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/models"
)

func (mr mongoRepository) GetNextNRaces(ctx context.Context, numRaces int) (models.Races, error) {
	nowTime := time.Now().UTC()
	nextRace := sort.Search(len(allRaces), func(i int) bool {
		return nowTime.Before(allRaces[i].CloseAt)
	})
	if nextRace == len(allRaces) {
		return nil, common.ErrNotFound
	}

	lastRace := nextRace + numRaces
	if lastRace > len(allRaces) {
		lastRace = len(allRaces)
	}
	return allRaces[nextRace:lastRace], nil
}

func (mr mongoRepository) GetRaceDetails(ctx context.Context, id string) (*models.RaceDetails, error) {
	raceDetails, ok := racesByID[id]
	if !ok {
		return nil, common.ErrNotFound
	}

	return raceDetails, nil
}
