package middleware

import (
	service "gokitdemo/services"

	"time"

	"github.com/go-kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.StringService
}

func (mw LoggingMiddleware) UpperCaseMW(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "UpperCase", "input", s, "output", output, "err", err, "took:", time.Since(begin))
	}(time.Now())
	output, err = mw.Next.MakeUpperCase(s)
	return
}

func (mw LoggingMiddleware) CountMW(s string) (output int) {
	defer func(begin time.Time) {
		mw.Logger.Log("method", "Count", "input", s, "output", output, "took:", time.Since(begin))
	}(time.Now())
	output = mw.Next.MakeCount(s)
	return
}

func test() string {
	return "can be imported"
}
