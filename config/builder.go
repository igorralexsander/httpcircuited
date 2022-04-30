package config

import (
	"time"
)

type ConfigBuilder struct {
	name               string
	baseUrl            string
	timeout            time.Duration
	defaultHeaders     map[string]string
	failureRatio       float64
	minRequests        uint32
	maxFailRequests    uint32
	delayPeriodOpened  time.Duration
	resetInterval      time.Duration
	requestsInHalfOpen uint32
}

func NewBuilder() *ConfigBuilder {

	return &ConfigBuilder{
		name:               "DEFAULT",
		timeout:            time.Duration(30000) * time.Millisecond,
		failureRatio:       0.6,
		minRequests:        3,
		maxFailRequests:    10,
		delayPeriodOpened:  time.Duration(30000) * time.Millisecond,
		resetInterval:      time.Duration(60000) * time.Millisecond,
		requestsInHalfOpen: 2,
		defaultHeaders:     make(map[string]string),
	}
}

func (b *ConfigBuilder) BaseUrl(url string) *ConfigBuilder {
	b.baseUrl = url
	return b
}

func (b *ConfigBuilder) Timeout(milliseconds int32) *ConfigBuilder {
	b.timeout = time.Duration(milliseconds) * time.Millisecond
	return b
}

func (b *ConfigBuilder) WithName(name string) *ConfigBuilder {
	b.name = name
	return b
}

func (b *ConfigBuilder) CircuitFailureRatio(ratio float64) *ConfigBuilder {
	b.failureRatio = ratio
	return b
}

func (b *ConfigBuilder) CircuitErrorsToOpen(errorCount int32) *ConfigBuilder {
	b.maxFailRequests = uint32(errorCount)
	return b
}

func (b *ConfigBuilder) CircuitDelayOpened(milliseconds int32) *ConfigBuilder {
	b.delayPeriodOpened = time.Duration(milliseconds) * time.Millisecond
	return b
}

func (b *ConfigBuilder) CircuitResetInterval(milliseconds int32) *ConfigBuilder {
	b.resetInterval = time.Duration(milliseconds) * time.Millisecond
	return b
}

func (b *ConfigBuilder) CircuitMinRequests(minRequests int32) *ConfigBuilder {
	b.minRequests = uint32(minRequests)
	return b
}

func (b *ConfigBuilder) CircuitRequestsInHalfOpen(count int32) *ConfigBuilder {
	b.requestsInHalfOpen = uint32(count)
	return b
}

func (b *ConfigBuilder) AddHeader(name string, value string) *ConfigBuilder {
	b.defaultHeaders[name] = value
	return b
}

func (b *ConfigBuilder) Build() *DownStreamConfig {
	return &DownStreamConfig{
		Name:               b.name,
		BaseUrl:            b.baseUrl,
		Timeout:            b.timeout,
		DefaultHeaders:     b.defaultHeaders,
		FailureRatio:       b.failureRatio,
		MinRequests:        b.minRequests,
		MaxFailRequests:    b.maxFailRequests,
		DelayPeriodOpened:  b.delayPeriodOpened,
		ResetInterval:      b.resetInterval,
		RequestsInHalfOpen: b.requestsInHalfOpen,
	}
}
