package controller

import (
	"context"

	"github.com/camotts/bggroup/controller/config"
	"github.com/camotts/bggroup/controller/store"
	"github.com/camotts/bggroup/rest_server_bggroup"
	"github.com/camotts/bggroup/rest_server_bggroup/operations"
	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	cfg *config.Config
	str *store.Store
)

func Run(inCfg *config.Config) error {
	cfg = inCfg

	swaggerSpec, err := loads.Embedded(rest_server_bggroup.SwaggerJSON, rest_server_bggroup.FlatSwaggerJSON)
	if err != nil {
		return errors.Wrap(err, "unable to load embedded swagger spec")
	}

	api := operations.NewBggroupAPI(swaggerSpec)
	api.AccountRegisterHandler = newRegisterHandler()

	if v, err := store.Open(cfg.Store); err == nil {
		str = v
	} else {
		return errors.Wrap(err, "unable to open store")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	if cfg.Maintenance != nil {
		if cfg.Maintenance.BGGSync != nil {
			go newMaintenanceBGGSyncAgent(ctx, cfg.Maintenance.BGGSync).run()
		}
		logrus.Infof("I should add maintenance upkeep: %v", ctx)
	}
	logrus.Infof("I should eventually do server stuff...")

	server := rest_server_bggroup.NewServer(api)
	server.Host = cfg.Endpoint.Host
	server.Port = cfg.Endpoint.Port

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		return errors.Wrap(err, "api server error")
	}

	return nil
}
