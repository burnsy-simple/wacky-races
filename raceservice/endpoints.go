package raceservice

import (
	"context"
	"net/url"
	"strings"

	"github.com/burnsy/wacky-races/models"
	"github.com/burnsy/wacky-races/payloads"
	"github.com/burnsy/wacky-races/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Endpoints are go-kit's abstractions around the API endpoints. They are typically all implemented on the server
// side but we usually only implement/expose a subset for the client side (as in other calling micro-services).
type Endpoints struct {
	GetRacesEndpoint endpoint.Endpoint
	GetRaceEndpoint  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetRacesEndpoint: MakeGetRacesEndpoint(s),
		GetRaceEndpoint:  MakeGetRaceEndpoint(s),
	}
}

// MakeClientEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the remote instance, via a transport/http.Client.
// Intended for use by clients of this service or even tests.
func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		GetRacesEndpoint: httptransport.NewClient("GET", tgt, encodeGetRacesRequest, decodeGetRacesResponse, options...).Endpoint(),
		GetRaceEndpoint:  httptransport.NewClient("GET", tgt, encodeGetRaceRequest, decodeGetRaceResponse, options...).Endpoint(),
	}, nil
}

// GetNextRaces retrieves the next N races that will be run
// swagger:route GET /races Races listRaces
//
// Lists the next N races.
//
// This will show the next 5 races by default.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       default: genericError
//       200: racesResponse
//       422: validationError
func (e Endpoints) GetNextRaces(ctx context.Context, numRaces int) (models.Races, error) {
	request := payloads.RacesReq{NumRaces: numRaces}
	response, err := e.GetRacesEndpoint(ctx, request)
	if err != nil {
		return models.Races{}, err
	}
	resp := response.(payloads.RacesResp)
	return resp.Races, resp.Err
}

// GetRaceDetails retrieves all the details the user needs to see for a given race
func (e Endpoints) GetRaceDetails(ctx context.Context, raceID string) (*models.RaceDetails, error) {
	request := payloads.RaceDetailsReq{RaceID: raceID}
	response, err := e.GetRaceEndpoint(ctx, request)
	if err != nil {
		return &models.RaceDetails{}, err
	}
	resp := response.(payloads.RaceDetailsResp)
	return resp.RaceDetails, resp.Err
}

// MakeGetRacesEndpoint returns an endpoint via the passed service.
func MakeGetRacesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(payloads.RacesReq)
		races, err := svc.GetNextRaces(ctx, req.NumRaces)
		return payloads.RacesResp{Races: races, Err: err}, nil
	}
}

// MakeGetRaceEndpoint returns an endpoint via the passed service.
func MakeGetRaceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(payloads.RaceDetailsReq)
		race, err := svc.GetRaceDetails(ctx, req.RaceID)
		return payloads.RaceDetailsResp{RaceDetails: race, Err: err}, nil
	}
}
