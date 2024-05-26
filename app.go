package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	transporthttp "github.com/go-kit/kit/transport/http"
)

type StringService interface {
	makeUpperCase(string) (string, error)
	makeCount(string) int
}

type StringStruct struct {
}

func (StringStruct) makeUpperCase(s string) (string, error) {
	fmt.Println("request with s:", s)
	if len(s) == 0 {
		return "", errors.New("empty string")
	}
	output := strings.ToUpper(s)
	fmt.Println("output with upper is:", output)
	return output, nil
}

func (StringStruct) makeCount(s string) int {
	return len(s)
}

//request and response in json

type upperCaseRequest struct {
	S string `json:"s"`
}

type upperCaseResponse struct {
	S     string `json:"s"`
	Error string `json:"error, omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	Count int `json:"count"`
}

func main() {
	fmt.Println("hello world")

	stringStruct := StringStruct{}

	upperCaseHandler := transporthttp.NewServer(makeUpperCaseEndpoint(stringStruct), decodeUpperCaseRequest, encodeResponse)
	countHandler := transporthttp.NewServer(makeCountEndpoint(stringStruct), decodeCountRequest, encodeResponse)

	http.Handle("/upper", upperCaseHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeUpperCaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request upperCaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	fmt.Println("upper case data in decoding:", request.S)
	return request, nil
}

func decodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func makeUpperCaseEndpoint(stringStruct StringStruct) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(upperCaseRequest)
		resp, err := stringStruct.makeUpperCase(req.S)
		if err != nil {
			fmt.Println("error occured in endpoint")
			return upperCaseResponse{"", err.Error()}, err
		}
		fmt.Println("response received at endpoint: ", resp)
		return upperCaseResponse{resp, ""}, nil
	}
}

func makeCountEndpoint(stringStruct StringStruct) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		countRequest := request.(countRequest)

		count := stringStruct.makeCount(countRequest.S)
		return countResponse{count}, nil
	}
}

//create handlers
