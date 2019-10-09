package user

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// decodeArithmeticRequest decode request params to struct
func decodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return UserRequest{
		RequestType: requestType,
		Uid:           a,
		B:           b,
	}, nil
}

// encodeArithmeticResponse encode response to return
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

