package downstream

import "time"

type DownStreamConfig struct {
	Name                string
	BaseUrl             string
	TimeoutMilliseconds time.Duration

	FailureRatio             float64
	MinRequests              uint32
	MaxFailRequests          uint32
	FailRequestsMilliseconds time.Duration
	RequestsInHalfOpen       uint32
}
