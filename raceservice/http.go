package raceservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/net/context"

	"github.com/burnsy/wacky-races/common"
	"github.com/burnsy/wacky-races/payloads"
	"github.com/burnsy/wacky-races/service"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(ctx context.Context, svc service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(svc)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("OPTIONS").Path("/races").HandlerFunc(preFlightHandler)
	r.Methods("GET").Path("/races").Handler(httptransport.NewServer(
		e.GetRacesEndpoint,
		decodeGetRacesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/race/{id}").Handler(httptransport.NewServer(
		e.GetRaceEndpoint,
		decodeGetRaceRequest,
		encodeResponse,
		options...,
	))
	return r
}

func preFlightHandler(w http.ResponseWriter, r *http.Request) {
	// REVISIT: Workaround for CORS when running locally

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://lvh.me:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "content-type,accept")
	w.WriteHeader(http.StatusOK)
}

func decodeGetRacesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := r.URL.Query()
	races, ok := vars["num_races"]
	numRaces := 0
	if !ok {
		numRaces = 5
	} else {
		numRaces, err = strconv.Atoi(races[0])
		if err != nil {
			return payloads.RacesReq{}, ErrBadData
		}
	}
	return payloads.RacesReq{NumRaces: numRaces}, nil
}

func decodeGetRaceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return payloads.RaceDetailsReq{
		RaceID: id,
	}, nil
}

func encodeGetRacesRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(payloads.RacesReq)
	req.Method, req.URL.Path = "GET", fmt.Sprintf("/races?numRaces=%d", r.NumRaces)
	return encodeRequest(ctx, req, request)
}

func encodeGetRaceRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(payloads.RaceDetailsReq)
	raceID := url.QueryEscape(r.RaceID)
	req.Method, req.URL.Path = "GET", "/race/"+raceID
	return encodeRequest(ctx, req, request)
}

func decodeGetRacesResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response payloads.RacesResp
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetRaceResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response payloads.RaceDetailsResp
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// encodeResponse is the common method to encode most/all response types.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Business-logic error - return as HTTP error.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Work-around for CORS when running locally
	w.Header().Set("Access-Control-Allow-Origin", "http://lvh.me:3000")
	return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

// errorer  to change the HTTP response code.
type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("developer error: encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case common.ErrNotFound:
		return http.StatusNotFound
	case ErrBadData:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
