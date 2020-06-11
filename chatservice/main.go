package main

import (
	"context"
	"flag"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	var logger log.Logger
	{

		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "chatService",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	flag.Pars()
	ctx := context.Background()
	var service chatService.ChatService
	{
		service = chatService.NewServer(logger)
	}
}
