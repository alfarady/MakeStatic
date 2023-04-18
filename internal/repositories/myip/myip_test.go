package myip_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/alfarady/makestatic/internal/repositories/myip"
	"github.com/alfarady/makestatic/internal/rest_client"
	"github.com/alfarady/makestatic/tests/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MyIPTestSuite struct {
	suite.Suite
	restClient *mocks.RestClient
	myip       myip.MyIPRepository
}

func (s *MyIPTestSuite) SetupTest() {
	s.restClient = &mocks.RestClient{}
	s.myip = myip.NewMyIPRepository(s.restClient)
}

func TestMyIPTest(t *testing.T) {
	suite.Run(t, new(MyIPTestSuite))
}

func (s *MyIPTestSuite) TestGet() {
	testcases := map[string]struct {
		ctx       context.Context
		resp      rest_client.RestResponse
		postError error
		wantErr   bool
	}{
		"success": {
			ctx: context.Background(),
			resp: rest_client.RestResponse{
				StatusCode: http.StatusOK,
				Body: `{
					"ip": "127.0.0.1"
				}`,
			},
		},
		"success not valid": {
			ctx: context.Background(),
			resp: rest_client.RestResponse{
				StatusCode: http.StatusOK,
				Body: `{
					"ip": "asede"
				}`,
			},
		},
		"error": {
			ctx:       context.Background(),
			postError: errors.New("failed"),
			wantErr:   true,
		},
		"error unmarshal": {
			ctx: context.Background(),
			resp: rest_client.RestResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"data": 500}`,
			},
			wantErr: true,
		},
		"error not 200": {
			ctx: context.Background(),
			resp: rest_client.RestResponse{
				StatusCode: http.StatusBadRequest,
				Body: `{
					"message": "failed"
				}`,
			},
		},
	}

	for _, tc := range testcases {
		if tc.resp.StatusCode != 0 || tc.postError != nil {
			s.restClient.On("Get", tc.ctx, mock.Anything).Return(tc.resp, tc.postError).Once()
		}

		res := s.myip.Get(tc.ctx)
		if tc.wantErr {
			s.Equal("", res)
		} else {
			s.Equal("127.0.0.1", res)
		}
	}
}
