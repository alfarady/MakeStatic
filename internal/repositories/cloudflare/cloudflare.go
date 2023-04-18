package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alfarady/makestatic/entity"
	"github.com/alfarady/makestatic/internal/rest_client"
)

type CloudflareRepository interface {
	UpdateDnsIP(ctx context.Context, params entity.CFUpdateDNSIp) error
}

type Cloudflare struct {
	client    rest_client.RestClient
	authToken string
}

func NewCloudflareRepository(client rest_client.RestClient, authToken string) *Cloudflare {
	return &Cloudflare{
		client:    client,
		authToken: authToken,
	}
}

func (r *Cloudflare) UpdateDnsIP(ctx context.Context, params entity.CFUpdateDNSIp) error {
	req := rest_client.RestRequest{
		Path: fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", params.ZoneID, params.RecordID),
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {fmt.Sprintf("Bearer %s", r.authToken)},
		},
		Payload: map[string]interface{}{"content": params.IP},
	}

	res, err := r.client.Post(ctx, req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
