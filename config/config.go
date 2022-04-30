package config

import "time"

type DownStreamConfig struct {
	Name               string
	BaseUrl            string
	Timeout            time.Duration
	DefaultHeaders     map[string]string
	FailureRatio       float64
	MinRequests        uint32
	MaxFailRequests    uint32
	DelayPeriodOpened  time.Duration
	ResetInterval      time.Duration
	RequestsInHalfOpen uint32
}
