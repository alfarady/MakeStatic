package rest_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
)

var ErrRequestTimeout = errors.New("timeout while waiting response")

type RestClient interface {
	Get(ctx context.Context, req RestRequest) (RestResponse, error)
	Post(ctx context.Context, req RestRequest) (RestResponse, error)
	Patch(ctx context.Context, req RestRequest) (RestResponse, error)
}

type RestRequest struct {
	Path    string
	Header  map[string][]string
	Payload interface{}
}

type RestResponse struct {
	Body       string
	StatusCode int
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type restClient struct {
	client HttpClient
}

func NewHttpClient(client HttpClient) *restClient {
	return &restClient{
		client: client,
	}
}

func (c *restClient) Get(ctx context.Context, req RestRequest) (RestResponse, error) {
	return c.doRequest(ctx, http.MethodGet, req)
}

func (c *restClient) Post(ctx context.Context, req RestRequest) (RestResponse, error) {
	return c.doRequest(ctx, http.MethodPost, req)
}

func (c *restClient) Patch(ctx context.Context, req RestRequest) (RestResponse, error) {
	return c.doRequest(ctx, http.MethodPatch, req)
}

func (c *restClient) doRequest(ctx context.Context, method string, restRequest RestRequest) (RestResponse, error) {
	var response RestResponse
	var body io.Reader
	if restRequest.Payload != nil {
		b, _ := json.Marshal(restRequest.Payload)
		body = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, restRequest.Path, body)
	if err != nil {
		return response, err
	}

	for key, value := range restRequest.Header {
		req.Header[key] = value
	}

	httpResponse, err := c.client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok {
			if netErr.Timeout() {
				return response, ErrRequestTimeout
			}
		}

		return response, err
	}
	defer httpResponse.Body.Close()

	responseBody, _ := io.ReadAll(httpResponse.Body)

	return RestResponse{
		Body:       string(responseBody),
		StatusCode: httpResponse.StatusCode,
	}, nil
}
