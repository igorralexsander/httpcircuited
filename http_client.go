package httpcircuited

import (
	"errors"
	"fmt"
	"github.com/igorralexsander/httpcircuited/config/downstream"
)

type HttpClient struct {
	downstreams map[string]Downstream
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c HttpClient) AddDownstream(config downstream.DownStreamConfig) {
	if _, exists := c.downstreams[config.Name]; !exists {
		service := NewDownstream(config)
		c.downstreams[config.Name] = *service
	}
}

func (c HttpClient) GetDownstream(name string) Downstream {
	if _, exists := c.downstreams[name]; !exists {
		errors.New(fmt.Sprintf("Downstream with name %s not found", name))
	}
	return c.downstreams[name]
}
