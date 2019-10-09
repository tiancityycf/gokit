package svc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	uid, ok := vars["uid"]
	if !ok {
		return nil, ErrorBadRequest
	}

	userId, _ := strconv.Atoi(uid)

	return UserRequest{
		Uid: userId,
	}, nil
}

func encodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Methods("GET").Path("/user/{uid}").Handler(kithttp.NewServer(
		endpoint,
		decodeUserRequest,
		encodeUserResponse,
		options...,
	))

	return r
}
