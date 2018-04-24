package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"context"

	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/middleware"
	"github.com/burnsy/wacky-races/raceservice"
	"github.com/burnsy/wacky-races/repository"
	"github.com/burnsy/wacky-races/service"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var repo repository.RaceRepository
	{
		repo = repository.NewRaceRepository(log.StandardLogger())
	}

	var svc service.Service
	{
		svc = raceservice.NewNextNService(repo, log.StandardLogger())
		svc = middleware.LoggingMiddleware(log.StandardLogger())(svc)
	}

	var ctx context.Context
	{
		ctx = context.Background()
		ctx = context.WithValue(ctx, common.LoggerKey, log.StandardLogger())
	}

	var h http.Handler
	{
		h = raceservice.MakeHTTPHandler(ctx, svc, log.StandardLogger())
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.WithFields(log.Fields{
			"transport": "HTTP",
			"addr":      *httpAddr,
		}).Info("Starting service")
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	log.Fatal("exit", <-errs)
}
