package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

type UserRequest struct {
	RequestType string `json:"request_type"`
	Uid           int    `json:"uid"`
	B           int    `json:"b"`
}

type UserResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

func MakeUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)

		var (
			res, uid int
			calError  error
		)

		uid = req.Uid

		if strings.EqualFold(req.RequestType, "Get") {
			res = svc.Get(uid)
		} else {
			return nil, nil
		}

		return UserResponse{Result: res, Error: calError}, nil
	}
}
