package raceservice

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/middleware"
	"github.com/burnsy/wacky-races/payloads"
	"github.com/burnsy/wacky-races/repository"
	"github.com/burnsy/wacky-races/service"
	log "github.com/sirupsen/logrus"
)

func TestGetNextRaces(t *testing.T) {
	h := createHandler()

	req := httptest.NewRequest(http.MethodGet, "/races", nil)
	recorder := httptest.NewRecorder()
	h.ServeHTTP(recorder, req)

	resp := recorder.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code: %d; got %d", 200, resp.StatusCode)
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		t.Errorf("Expected content type to be %s; got %s", "application/json", resp.Header.Get("Content-Type"))
	}

	if len(body) <= 2 {
		t.Errorf("Expected non-empty content; got %s", string(body))
	}

	var payload payloads.RacesResp
	err := json.Unmarshal(body, &payload)
	if err != nil {
		t.Fatalf("Expected races; got error %v; body = %v", err, string(body))
	}

	if len(payload.Races) != defaultNumRaces {
		t.Errorf("Expected 5 races; got %d", len(payload.Races))
	}
}

func createHandler() http.Handler {
	var repo repository.RaceRepository
	{
		repo = repository.NewRaceRepository(log.StandardLogger())
	}

	var svc service.Service
	{
		svc = NewNextNService(repo, log.StandardLogger())
		svc = middleware.LoggingMiddleware(log.StandardLogger())(svc)
	}

	var ctx context.Context
	{
		ctx = context.Background()
		ctx = context.WithValue(ctx, common.LoggerKey, log.StandardLogger())
	}

	var h http.Handler
	{
		h = MakeHTTPHandler(ctx, svc, log.StandardLogger())
	}

	return h
}
