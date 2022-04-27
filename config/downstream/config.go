package downstream

import "time"

type DownStreamConfig struct {
	Name                     string
	BaseUrl                  string
	TimeoutMilliseconds      time.Duration
	DefaultHeaders           map[string]string
	FailureRatio             float64
	MinRequests              uint32
	MaxFailRequests          uint32
	FailRequestsMilliseconds time.Duration
	RequestsInHalfOpen       uint32
}

func (m *DownStreamConfig) AddDefaultHeader(name string, value string) *DownStreamConfig {
	if m.DefaultHeaders == nil {
		m.DefaultHeaders = make(map[string]string)
	}
	m.DefaultHeaders[name] = value
	return m
}
