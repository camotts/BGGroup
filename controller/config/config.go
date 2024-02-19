package config

import (
	"time"

	"github.com/camotts/bggroup/controller/env"
	"github.com/camotts/bggroup/controller/store"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
)

type Config struct {
	Endpoint    *EndpointConfig
	Maintenance *MaintenanceConfig
	Store       *store.Config
}

type EndpointConfig struct {
	Host string
	Port int
}

type MaintenanceConfig struct {
	BGGSync *BGGSyncMaintenanceConfig
}

type BGGSyncMaintenanceConfig struct {
	Frequency time.Duration
}

func DefaultConfig() *Config {
	return &Config{}
}

func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()
	if err := cf.BindYaml(cfg, path, env.GetCfOptions()); err != nil {
		return nil, errors.Wrapf(err, "unable to load controller config %v", path)
	}
	return cfg, nil
}
