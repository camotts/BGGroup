package controller

import (
	"context"
	"time"

	"github.com/camotts/bggroup/controller/config"
	"github.com/fzerorubigd/gobgg"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
)

type maintenanceBGGSyncAgent struct {
	cfg *config.BGGSyncMaintenanceConfig
	ctx context.Context
}

func newMaintenanceBGGSyncAgent(ctx context.Context, cfg *config.BGGSyncMaintenanceConfig) *maintenanceBGGSyncAgent {
	return &maintenanceBGGSyncAgent{
		cfg: cfg,
		ctx: ctx,
	}
}

func (ma *maintenanceBGGSyncAgent) run() {
	logger := logrus.WithField("Type", "Agent").WithField("Name", "BGGSync")
	logger.Info("started")
	defer logger.Info("exited")

	rl := ratelimit.New(10, ratelimit.Per(60*time.Second))
	_ = gobgg.NewBGGClient(gobgg.SetLimiter(rl))

	ticker := time.NewTicker(ma.cfg.Frequency)
	for {
		select {
		case <-ma.ctx.Done():
			{
				ticker.Stop()
				return
			}
		case <-ticker.C:
			logger.Info("Time for a BGG Sync!")
		}
	}
}
