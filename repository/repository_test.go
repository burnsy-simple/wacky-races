package repository

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetRaces(t *testing.T) {
	logger := log.StandardLogger()
	repo := NewRaceRepository(logger)

	races, err := repo.GetNextNRaces(nil, 3)
	if err != nil {
		t.Errorf("received error getting races: %v", err)
	}
	if len(races) != 3 {
		t.Errorf("expected 3 races; got %d", len(races))
	}
}

func BenchmarkGetRaces(b *testing.B) {
	logger := log.StandardLogger()
	repo := NewRaceRepository(logger)

	for i := 0; i < b.N; i++ {
		races, _ := repo.GetNextNRaces(nil, 5)
		if len(races) != 5 {
			b.Logf("expected 5 races; got %d", len(races))
		}
	}
}
