package transport

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	service "gokitdemo/services"

	"github.com/go-kit/kit/endpoint"
)

//request and response in json

type UpperCaseRequest struct {
	S string `json:"s"`
}

type UpperCaseResponse struct {
	S     string `json:"s"`
	Error string `json:"error, omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	Count int `json:"count"`
}

func hello() string {
	return "hello world"
}

func DecodeUpperCaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpperCaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	fmt.Println("upper case data in decoding:", request.S)
	return request, nil
}

func DecodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MakeUpperCaseEndpoint(stringStruct service.StringStruct) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(UpperCaseRequest)
		resp, err := stringStruct.MakeUpperCase(req.S)
		if err != nil {
			fmt.Println("error occured in endpoint")
			return UpperCaseResponse{"", err.Error()}, err
		}
		fmt.Println("response received at endpoint: ", resp)
		return UpperCaseResponse{resp, ""}, nil
	}
}

func MakeCountEndpoint(stringStruct service.StringStruct) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		countRequest := request.(CountRequest)

		count := stringStruct.MakeCount(countRequest.S)
		return CountResponse{count}, nil
	}
}
