// Package client provides a raceService client based on a predefined Consul
// service name and relevant tags.
package client

import (
	"io"
	"time"

	"github.com/burnsy/wacky-races/raceService"
	consulapi "github.com/hashicorp/consul/api"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
)

// New returns a service that's load-balanced over instances of raceService found
// in the provided Consul server. The mechanism of looking up raceService
// instances in Consul is hard-coded into the client.
func New(consulAddr string, logger log.Logger) (raceService.Service, error) {
	apiclient, err := consulapi.NewClient(&consulapi.Config{
		Address: consulAddr,
	})
	if err != nil {
		return nil, err
	}

	// As the implementer of raceService, we declare and enforce these
	// parameters for all of the raceService consumers.
	var (
		consulService = "raceService"
		consulTags    = []string{"prod"}
		passingOnly   = true
		retryMax      = 3
		retryTimeout  = 500 * time.Millisecond
	)

	var (
		sdclient  = consul.NewClient(apiclient)
		instancer = consul.NewInstancer(sdclient, logger, consulService, consulTags, passingOnly)
		endpoints raceService.Endpoints
	)
	{
		factory := factoryFor(raceService.MakeGetRacesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetRacesEndpoint = retry
	}
	{
		factory := factoryFor(raceService.MakeGetRaceEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetRaceEndpoint = retry
	}

	return endpoints, nil
}

func factoryFor(makeEndpoint func(raceService.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		service, err := raceService.MakeClientEndpoints(instance)
		if err != nil {
			return nil, nil, err
		}
		return makeEndpoint(service), nil, nil
	}
}
