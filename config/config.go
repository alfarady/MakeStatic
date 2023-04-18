package config

import "github.com/joeshaw/envdecode"

type Config struct {
	AppName       string `env:"APP_NAME,default=MakeStatic"`
	CheckInterval int    `env:"CHECK_INTERVAL,default=10"`

	RestClientOption
	Cloudflare
}

type RestClientOption struct {
	MaxIdleConns        int `env:"REST_CLIENT_MAX_IDLE_CONNECTION"`
	MaxIdleConnsPerHost int `env:"REST_CLIENT_MAX_IDLE_CONNECTION_PER_HOST"`
	MaxConnsPerHost     int `env:"REST_CLIENT_MAX_CONNECTION_PER_HOST"`
	IdleConnTimeoutMs   int `env:"REST_IDLE_CONNECTION_TIMEOUT"`
	TimeoutMs           int `env:"REST_CLIENT_TIMEOUT"`
}

type Cloudflare struct {
	AuthToken string   `env:"CF_AUTH_TOKEN,required"`
	ZoneID    string   `env:"CF_ZONE_ID,required"`
	RecordIds []string `env:"CF_RECORD_IDS,required"`
}

// NewConfig initialize new config
func NewConfig() *Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
