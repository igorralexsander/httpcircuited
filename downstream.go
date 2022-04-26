package httpclient

import (
	"github.com/go-resty/resty/v2"
	"httpclient/circuitbreaker"
	"httpclient/config/downstream"
)

type Downstream struct {
	client *resty.Client
	cb     circuitbreaker.CircuitBreaker
}

func NewDownstream(config downstream.DownStreamConfig) *Downstream {
	restyClient := resty.New()
	restyClient.SetTimeout(config.TimeoutMilliseconds)
	restyClient.SetBaseURL(config.BaseUrl)

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