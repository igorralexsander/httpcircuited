package circuitbreaker

import (
	"github.com/igorralexsander/httpcircuited/config"
	"github.com/sony/gobreaker"
)

func MakeCircuitBreaker(config config.DownStreamConfig) CircuitBreaker {
	cb := CircuitBreaker{
		minRequests:           config.MinRequests,
		failureRatioThreshold: config.FailureRatio,
	}

	cb.cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          "cb-" + config.Name,
		MaxRequests:   config.RequestsInHalfOpen,
		Interval:      config.DelayPeriodOpened,
		Timeout:       config.DelayPeriodOpened,
		ReadyToTrip:   cb.ratioThreshold,
		OnStateChange: cb.watcher,
	})
	return cb
}
