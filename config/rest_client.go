package config

import (
	"net/http"
	"time"

	"github.com/alfarady/makestatic/internal/rest_client"
)

func NewRestClient(cfg RestClientOption) rest_client.RestClient {
	return rest_client.NewHttpClient(cfg.NewHttpClient())
}

func (c RestClientOption) NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        c.MaxIdleConns,
			MaxIdleConnsPerHost: c.MaxIdleConnsPerHost,
			MaxConnsPerHost:     c.MaxConnsPerHost,
			IdleConnTimeout:     time.Duration(c.IdleConnTimeoutMs) * time.Millisecond,
		},
		Timeout: time.Duration(c.TimeoutMs) * time.Millisecond,
	}
}
