package httpcircuited

import (
	"github.com/go-resty/resty/v2"
	"github.com/igorralexsander/httpcircuited/circuitbreaker"
	"github.com/igorralexsander/httpcircuited/config"
	"time"
)

type Downstream struct {
	client *resty.Client
	cb     circuitbreaker.CircuitBreaker
}

func NewDownstream(config config.DownStreamConfig) *Downstream {
	restyClient := resty.New()
	restyClient.SetTimeout(config.TimeoutMilliseconds * time.Millisecond)
	restyClient.SetBaseURL(config.BaseUrl)

	if config.DefaultHeaders != nil {
		for key, value := range config.DefaultHeaders {
			restyClient.SetHeader(key, value)
		}
	}

	circuitBreaker := circuitbreaker.MakeCircuitBreaker(config)
	return &Downstream{
		client: restyClient,
		cb:     circuitBreaker,
	}
}

func (d Downstream) Get(path string) ([]byte, error) {
	response, err := d.cb.Fetch(func() (interface{}, error) {
		return d.client.R().Get(path)
	})
	if err != nil {
		return nil, err
	}
	return response.(*resty.Response).Body(), nil
}

func (d Downstream) Post(path string, body interface{}) ([]byte, error) {
	response, err := d.cb.Fetch(func() (interface{}, error) {
		return d.client.R().SetBody(body).Post(path)
	})
	if err != nil {
		return nil, err
	}
	return response.(*resty.Response).Body(), nil
}
