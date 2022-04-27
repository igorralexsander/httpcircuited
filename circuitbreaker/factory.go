package circuitbreaker

import (
	"github.com/igorralexsander/httpcircuited/config"
	"github.com/sony/gobreaker"
	"time"
)

func MakeCircuitBreaker(config config.DownStreamConfig) CircuitBreaker {
	cb := CircuitBreaker{
		minRequests:           config.MinRequests,
		failureRatioThreshold: config.FailureRatio,
	}

	cb.cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          "cb-" + config.Name,
		MaxRequests:   config.RequestsInHalfOpen,
		Interval:      time.Millisecond * config.FailRequestsMilliseconds,
		Timeout:       time.Millisecond * config.TimeoutMilliseconds,
		ReadyToTrip:   cb.ratioThreshold,
		OnStateChange: cb.watcher,
	})
	return cb
}
