package circuitbreaker

import (
	"fmt"
	"github.com/sony/gobreaker"
)

type CircuitBreaker struct {
	failRequests          uint32
	failureRatioThreshold float64
	cb                    *gobreaker.CircuitBreaker
}

func (c CircuitBreaker) Fetch(request func() (interface{}, error)) (interface{}, error) {
	response, err := c.cb.Execute(request)
	return response, err
}

func (c *CircuitBreaker) ratioThreshold(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= c.failRequests && failureRatio >= c.failureRatioThreshold
}

func (c *CircuitBreaker) watcher(name string, from gobreaker.State, to gobreaker.State) {
	if to == gobreaker.StateOpen {
		fmt.Println(fmt.Sprintf("[CIRCUIT BREAKER] The circuit %s State Change from %s to %s", name, from, to))
	} else {
		fmt.Println(fmt.Sprintf("[CIRCUIT BREAKER] The circuit %s State Change from %s to %s", name, from, to))
	}
}
