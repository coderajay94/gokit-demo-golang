package main

import (
	"fmt"
	"log"
	"net/http"

	service "gokitdemo/services"
	transport "gokitdemo/transports"

	"github.com/go-kit/kit/endpoint"
	transporthttp "github.com/go-kit/kit/transport/http"
)

func main() {
	fmt.Println("welcome to app")

	//uppercase =

	svc := service.StringStruct{}

	var uppercaseEndpoint endpoint.Endpoint
	uppercaseEndpoint = transport.MakeUpperCaseEndpoint(svc)

	var countEndpoint endpoint.Endpoint
	countEndpoint = transport.MakeCountEndpoint(svc)

	uppercaseHandler := transporthttp.NewServer(uppercaseEndpoint, transport.DecodeUpperCaseRequest, transporthttp.EncodeJSONResponse)
	countHandler := transporthttp.NewServer(countEndpoint, transport.DecodeCountRequest, transport.EncodeResponse)

	http.Handle("/upper", uppercaseHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
