package middleware

import "github.com/go-kit/log"

type LoggingMiddleware struct {
	Logger log.Logger
}
