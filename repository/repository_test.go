package repository

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestGetRaces(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	repo := NewRaceRepository(logger)

	races, err := repo.GetNextNRaces(nil, 3)
	if err != nil {
		t.Errorf("received error getting races: %v", err)
	}
	if len(races) != 3 {
		t.Errorf("expected 3 races; got %d", len(races))
	}
}
