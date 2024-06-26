package main

import (
	"context"
	"phanes/utils"

	"phanes/assistant"
	"phanes/bll"
	"phanes/client"
	"phanes/collector"
	"phanes/config"
	"phanes/event"
	"phanes/server"
	"phanes/store"
)

type InitFunc func() func()

func main() {
	var (
		// system init func
		bootstraps = []InitFunc{
			config.Init,
			collector.Init,
			server.Init,
			client.Init,
			store.Init,
			assistant.Init,
			bll.Init,
		}
	)

	event.Init()
	for _, bootstrap := range bootstraps {
		cancel := bootstrap()
		event.Subscribe(event.EventExit, func(ctx context.Context, e event.Event, payload interface{}) {
			cancel()
		})
	}

	utils.Throw(config.MicroService.Run())
}
