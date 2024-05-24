package transports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coderajay94/gokit-demo-golang/services"
	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	S   string `json:"s"`
	Err string `json:"err, omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	Count int    `json:"count"`
	Err   string `json:"err, omitempty"`
}

func makeUpperCaseEndpoint(ser services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)

		res, err := ser.UpperCase(req.S)
		if err != nil {
			return uppercaseResponse{res, err.Error()}, err
		} else {
			return uppercaseResponse{res, ""}, nil
		}
	}
}

func makeCountEndpoint(ser services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		count := ser.Count(req.S)
		return countResponse{count, ""}, nil
	}
}

func decodeUpperCaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request countRequest

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}
