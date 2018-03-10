package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/middleware"
	"github.com/burnsy/wacky-races/raceservice"
	"github.com/burnsy/wacky-races/repository"
	"github.com/go-kit/kit/log"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

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

	var svc raceservice.Service
	{
		svc = raceservice.NewNextNService(repo, logger)
		svc = middleware.LoggingMiddleware(logger)(svc)
	}

	var ctx context.Context
	{
		ctx = context.Background()
		ctx = context.WithValue(ctx, common.LoggerKey, log.With(logger, "component", common.LoggerKey))
	}

	var h http.Handler
	{
		h = raceservice.MakeHTTPHandler(ctx, svc, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
