package raceservice

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/middleware"
	"github.com/burnsy/wacky-races/repository"
	"github.com/burnsy/wacky-races/service"
	"github.com/go-kit/kit/log"
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
}

func createHandler() http.Handler {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var repo repository.RaceRepository
	{
		repo = repository.NewRaceRepository(log.With(logger, "component", common.RepositoryKey))
	}

	var svc service.Service
	{
		svc = NewNextNService(repo, logger)
		svc = middleware.LoggingMiddleware(logger)(svc)
	}

	var ctx context.Context
	{
		ctx = context.Background()
		ctx = context.WithValue(ctx, common.LoggerKey, log.With(logger, "component", common.LoggerKey))
	}

	var h http.Handler
	{
		h = MakeHTTPHandler(ctx, svc, log.With(logger, "component", "HTTP"))
	}

	return h
}
