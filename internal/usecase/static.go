package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/alfarady/makestatic/entity"
	"github.com/alfarady/makestatic/internal/repositories/cloudflare"
	"github.com/alfarady/makestatic/internal/repositories/myip"
)

type StaticProvider struct {
	cloudflare cloudflare.CloudflareRepository
	myip       myip.MyIPRepository
}

type StaticUsecase interface {
	MakeStatic(ctx context.Context, params entity.StaticParams) error
}

func NewStaticUsecase(cloudflare cloudflare.CloudflareRepository, myip myip.MyIPRepository) *StaticProvider {
	return &StaticProvider{cloudflare, myip}
}

func (u *StaticProvider) MakeStatic(ctx context.Context, params entity.StaticParams) (*entity.StaticParams, error) {
	log.Println("[MakeStatic] Checking ip changes...")
	currIp := ""

	for i := 0; i < 3; i++ {
		currIp = u.myip.Get(ctx)
		if currIp != "" {
			break
		}
		log.Printf("[MakeStatic] Failed to fetch current ip: %s\n\n", currIp)
		time.Sleep(time.Second * time.Duration(math.Pow(float64(i), 2)))
	}
	if currIp == "" {
		log.Printf("[MakeStatic] Failed to fetch current ip\n\n")
		return &params, errors.New("failed to fetch ip")
	}

	if currIp == params.PrevIP {
		log.Printf("[MakeStatic] No ip changes\n\n")
		return &params, nil
	}

	log.Printf("[MakeStatic] Success to fetch current ip: %s\n", currIp)

	params.PrevIP = currIp
	for _, id := range params.RecordIds {
		log.Printf("[MakeStatic] Updating DNS Record %s...\n", id)
		err := u.cloudflare.UpdateDnsIP(ctx, entity.CFUpdateDNSIp{
			RecordID: id,
			ZoneID:   params.ZoneID,
			IP:       currIp,
		})
		if err != nil {
			fmt.Printf("[MakeStatic] Failed to update DNS Record %s\n", id)
			continue
		}
		log.Printf("[MakeStatic] Success to update DNS Record %s\n", id)
	}

	log.Printf("[MakeStatic] Finish\n\n")
	return &params, nil
}
