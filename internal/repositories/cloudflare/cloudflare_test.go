package cloudflare_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/alfarady/makestatic/entity"
	"github.com/alfarady/makestatic/internal/repositories/cloudflare"
	"github.com/alfarady/makestatic/internal/rest_client"
	"github.com/alfarady/makestatic/tests/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CloudflareTestSuite struct {
	suite.Suite
	restClient *mocks.RestClient
	cloudflare cloudflare.CloudflareRepository
}

func (s *CloudflareTestSuite) SetupTest() {
	s.restClient = &mocks.RestClient{}
	s.cloudflare = cloudflare.NewCloudflareRepository(s.restClient, "test")
}

func TestCloudflareTest(t *testing.T) {
	suite.Run(t, new(CloudflareTestSuite))
}

func (s *CloudflareTestSuite) TestUpdateDnsIP() {
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
					"message": "success"
				}`,
			},
		},
		"error": {
			ctx:       context.Background(),
			postError: errors.New("failed"),
			wantErr:   true,
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

	for name, tc := range testcases {
		if tc.resp.StatusCode != 0 || tc.postError != nil {
			s.restClient.On("Post", tc.ctx, mock.Anything).Return(tc.resp, tc.postError).Once()
		}

		err := s.cloudflare.UpdateDnsIP(tc.ctx, entity.CFUpdateDNSIp{})
		if tc.wantErr {
			s.Error(err, name)
		} else {
			s.NoError(err, name)
		}
	}
}
