package myip

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/alfarady/makestatic/entity"
	"github.com/alfarady/makestatic/internal/rest_client"
)

type MyIPRepository interface {
	Get(ctx context.Context) string
}

type MyIP struct {
	client rest_client.RestClient
}

func NewMyIPRepository(client rest_client.RestClient) *MyIP {
	return &MyIP{
		client: client,
	}
}

func (r *MyIP) Get(ctx context.Context) string {
	req := rest_client.RestRequest{
		Path: "https://api.myip.com/",
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	res, err := r.client.Get(ctx, req)
	if err != nil {
		return ""
	}

	if res.StatusCode != http.StatusOK {
		return ""
	}

	var data entity.MyIPResponse
	if err := json.Unmarshal([]byte(res.Body), &data); err != nil {
		return ""
	}

	if !isValidIpAddr(data.IP) {
		return ""
	}

	return data.IP
}

func isValidIpAddr(ip string) bool {
	return net.ParseIP(ip) != nil
}
