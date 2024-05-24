package services

import (
	"strings"
	"time"

	"github.com/go-kit/log"
)

type StringService interface {
	UpperCase(st string) (string, error)
	Count(st string) int
}

type stringservice struct{}

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) UpperCase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "UpperCase", "input", s, "output", output, "err", err, "time elapsed", time.Since(begin))
	}(time.Now())

	output, err = mw.next.UpperCase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (output int) {

	defer func(begin time.Time) {
		mw.logger.Log("method", "Count", "input", s, "output", output, "time elapsed", time.Since(begin))
	}(time.Now())

	output = mw.next.Count(s)
	return
}

func (s stringservice) UpperCase(st string) (string, error) {

	if len(st) == 0 {
		return "", nil
	} else {
		return strings.ToUpper(st), nil
	}
}

func (s stringservice) Count(st string) int {
	return len(st)
}
