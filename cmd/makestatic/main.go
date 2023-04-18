package main

import (
	"context"
	"log"
	"time"

	"github.com/alfarady/makestatic/config"
	"github.com/alfarady/makestatic/entity"
	"github.com/alfarady/makestatic/internal/repositories/cloudflare"
	"github.com/alfarady/makestatic/internal/repositories/myip"
	"github.com/alfarady/makestatic/internal/usecase"
	"github.com/go-co-op/gocron"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	cfg := config.NewConfig()

	log.Printf("[MakeStatic] Listening to ip changes and update DNS Record ZoneID: %s\n\n", cfg.Cloudflare.ZoneID)

	ctx := context.Background()
	restClient := config.NewRestClient(cfg.RestClientOption)
	s := gocron.NewScheduler(time.Local)

	cloudflareRepository := cloudflare.NewCloudflareRepository(restClient, cfg.AuthToken)
	myIpRepository := myip.NewMyIPRepository(restClient)

	staticUsecase := usecase.NewStaticUsecase(cloudflareRepository, myIpRepository)

	prevIp := ""
	s.Every(cfg.CheckInterval).Minutes().Do(func() {
		res, _ := staticUsecase.MakeStatic(ctx, entity.StaticParams{
			PrevIP:    prevIp,
			RecordIds: cfg.RecordIds,
			ZoneID:    cfg.ZoneID,
		})
		prevIp = res.PrevIP
	})

	s.StartAsync()
	s.StartBlocking()
}
