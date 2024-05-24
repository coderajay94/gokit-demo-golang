package main

import (
	"fmt"
	"os"

	"github.com/coderajay94/gokit-demo-golang/services"
	"github.com/go-kit/log"
)

func main() {
	fmt.Println("starting go-kit demo using sepration of concern....")

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc services.StringService
	//svc = services.StringService{}
	//svc = loggingMiddleware{logger, svc}

}
