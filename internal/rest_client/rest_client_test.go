package rest_client_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alfarady/makestatic/internal/rest_client"
	"github.com/alfarady/makestatic/tests/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HttpClientSuite struct {
	suite.Suite

	ctx          context.Context
	restClient   rest_client.RestClient
	mockClient   *mocks.HttpClient
	httpResponse *http.Response
	restRequest  rest_client.RestRequest
	restResponse rest_client.RestResponse
}

func TestHttpClientSuite(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}

type mockNetError struct {
	error
	errorTimeout bool
	errorTemp    bool
}

func (e mockNetError) Timeout() bool {
	return e.errorTimeout
}

func (e mockNetError) Temporary() bool {
	return e.errorTemp
}

func (s *HttpClientSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockClient = new(mocks.HttpClient)
	s.restClient = rest_client.NewHttpClient(s.mockClient)
	s.httpResponse = &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{ "message" : "ok"  }`)),
		StatusCode: 200,
	}
	s.restResponse = rest_client.RestResponse{
		Body:       `{ "message" : "ok"  }`,
		StatusCode: 200,
	}
	s.restRequest = rest_client.RestRequest{
		Path:   "some/url",
		Header: map[string][]string{"header": {"value"}},
		Payload: struct {
			Key string `json:"key"`
		}{
			Key: "value",
		},
	}
}

func (s *HttpClientSuite) Test_Get_Success() {
	s.mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(s.httpResponse, nil)
	res, err := s.restClient.Get(s.ctx, s.restRequest)

	s.Assert().Nil(err)
	s.Assert().Equal(s.restResponse, res)
}

func (s *HttpClientSuite) Test_Get_ErrorCreateRequest() {
	s.restRequest.Path = "/some/ur%%l"
	res, err := s.restClient.Get(s.ctx, s.restRequest)

	s.Assert().Error(err)
	s.Assert().Zero(res)
}

func (s *HttpClientSuite) Test_Get_ErrorDoRequest() {
	expectedError := errors.New("request error")

	s.mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, expectedError)

	res, err := s.restClient.Get(s.ctx, s.restRequest)
	s.Assert().ErrorIs(expectedError, err)
	s.Assert().Zero(res)
}

func (s *HttpClientSuite) Test_Get_Timeout() {
	expectedError := mockNetError{
		errorTimeout: true,
	}

	s.mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, expectedError)

	res, err := s.restClient.Get(s.ctx, s.restRequest)
	s.Assert().ErrorIs(rest_client.ErrRequestTimeout, err)
	s.Assert().Zero(res)
}

func (s *HttpClientSuite) Test_Post_Success() {
	s.mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(s.httpResponse, nil)
	res, err := s.restClient.Post(s.ctx, s.restRequest)

	s.Assert().Nil(err)
	s.Assert().Equal(s.restResponse, res)
}

func (s *HttpClientSuite) Test_Patch_Success() {
	s.mockClient.On("Do", mock.AnythingOfType("*http.Request")).Return(s.httpResponse, nil)
	res, err := s.restClient.Patch(s.ctx, s.restRequest)

	s.Assert().Nil(err)
	s.Assert().Equal(s.restResponse, res)
}
