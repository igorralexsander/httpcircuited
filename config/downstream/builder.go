package downstream

import "time"

type ConfigBuilder struct {
	name                string
	BaseUrl             string
	TimeoutMilliseconds time.Duration
	DefaultHeaders      map[string]string
	FailureRatio        float64
	MinRequests         uint32
	MaxFailRequests     uint32
	DelayPeriodRetry    time.Duration
	RequestsInHalfOpen  uint32
}

func NewBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		name:                "DEFAULT",
		TimeoutMilliseconds: time.Duration(30000),
		FailureRatio:        0.6,
		MinRequests:         3,
		MaxFailRequests:     10,
		DelayPeriodRetry:    time.Duration(15000),
		RequestsInHalfOpen:  2,
	}
}

func (b *ConfigBuilder) WithName(name string) *ConfigBuilder {
	b.name = name
	return b
}

func (b *ConfigBuilder) Build() *DownStreamConfig {
	return &DownStreamConfig{
		Name:               b.name,
		BaseUrl:            b.BaseUrl,
		Timeout:            b.TimeoutMilliseconds,
		DefaultHeaders:     b.DefaultHeaders,
		FailureRatio:       b.FailureRatio,
		MinRequests:        b.MinRequests,
		MaxFailRequests:    b.MaxFailRequests,
		DelayPeriodRetry:   b.DelayPeriodRetry,
		RequestsInHalfOpen: b.RequestsInHalfOpen,
	}
}
