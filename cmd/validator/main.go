package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/manifold-inc/targon/internal/validator"
	"github.com/manifold-inc/targon/internal/validator/callbacks"
	"github.com/manifold-inc/targon/internal/validator/setup"

	"github.com/subtrahend-labs/gobt/boilerplate"
)

func main() {
	deps := setup.Init()
	deps.Log.Infof(
		"Starting validator with key [%s] on chain [%s] version [%d]",
		deps.Hotkey.Address,
		deps.Env.ChainEndpoint,
		deps.Env.Version,
	)
	if deps.Mongo != nil {
		defer func() {
			if err := deps.Mongo.Disconnect(context.Background()); err != nil {
				deps.Log.Errorw("failed disconnecting from mongo", "error", err)
			}
		}()
	}

	core := validator.CreateCore(deps)
	v := boilerplate.NewChainSubscriber()
	deps.Log.Infof("Creating validator on netuid [%d]", deps.Env.Netuid)

	callbacks.AddBlockCallbacks(v, core)

	v.SetOnSubscriptionError(func(e error) {
		deps.Log.Errorw("Subscription Error", "error", e)
	})
	err := validator.LoadMongoBackup(core)
	if err != nil {
		core.Deps.Log.Warn("Failed to load last checkpoint")
	}
	if err == nil {
		core.Deps.Log.Info("Loaded checkpoint from mongo")
	}
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		v.Stop()
	}()

	for {
		err := v.Start(deps.Client)
		if err != nil {
			deps.Log.Errorw("Subscription Error", "error", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	core.Deps.Log.Info("Shutting down validator")
	err = validator.SaveMongoBackup(core)
	if err != nil {
		core.Deps.Log.Errorw("Failed saving backup of state", "error", err)
	}
}
